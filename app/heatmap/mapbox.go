package heatmap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"log"
	"net/http"
	"os"
)

const _mapboxUploadsBaseURL = "https://api.mapbox.com/uploads/v1/"

var (
	mapboxUploadAccessToken = os.Getenv("MAPBOX_UPLOAD_ACCESS_TOKEN")
	mapBoxUsername          = os.Getenv("MAPBOX_USERNAME")
	mapboxUploadsBaseURL    = fmt.Sprintf("%s%s/", _mapboxUploadsBaseURL, mapBoxUsername)
)

type credentialsResponse struct {
	AccessKeyId     string `json:"accessKeyId"`
	Bucket          string `json:"bucket"`
	Key             string `json:"key"`
	SecretAccessKey string `json:"secretAccessKey"`
	SessionToken    string `json:"sessionToken"`
	URL             string `json:"url"`
}

type createUploadBody struct {
	URL     string `json:"url"`
	Tileset string `json:"tileset"`
}

type createUploadResponse struct {
	Complete bool   `json:"complete"`
	Tileset  string `json:"tileset"`
	Error    string `json:"error"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Modified string `json:"modified"`
	Created  string `json:"created"`
	Owner    string `json:"owner"`
	Progress int    `json:"progress"`
}

func updateMapbox(ds data.Store, filename string) error {
	log.Print("Uploading heatmap to Mapbox")
	s3Url, err := uploadHeatmap(filename)
	if err != nil {
		return err
	}

	log.Print("Creating upload in Mapbox")
	err = createUpload(ds, s3Url)
	if err != nil {
		return err
	}

	return nil
}

func uploadHeatmap(fileName string) (string, error) {
	credentialsUrl := fmt.Sprintf("%scredentials?access_token=%s", mapboxUploadsBaseURL, mapboxUploadAccessToken)

	log.Printf("Getting credentials from %s", credentialsUrl)
	resp, err := http.Get(credentialsUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	s3Credentials := credentialsResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&s3Credentials); err != nil {
		return "", err
	}
	log.Printf("Got credentials: %v", credentialsResponse{})

	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				s3Credentials.AccessKeyId,
				s3Credentials.SecretAccessKey,
				s3Credentials.SessionToken,
			),
		),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return "", err
	}

	s3Client := s3.NewFromConfig(awsConfig)

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	log.Printf("Uploading file: %s to bucket: %s", fileName, s3Credentials.Bucket)
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Credentials.Bucket),
		Key:    aws.String(s3Credentials.Key),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return s3Credentials.URL, nil
}

func createUpload(ds data.Store, s3Url string) error {
	tileSetID, err := TileSetID(ds)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?access_token=%s", mapboxUploadsBaseURL, mapboxUploadAccessToken)
	body := createUploadBody{
		URL:     s3Url,
		Tileset: tileSetID,
	}
	uploadData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	fmt.Printf("Creating upload with body: %s\n", uploadData)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(uploadData))
	if err != nil {
		return err
	}
	fmt.Printf("Response code: %d\n", resp.StatusCode)
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}
	response := createUploadResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	responseData, _ := json.Marshal(response)
	log.Printf("Got response: %s", responseData)

	return nil
}

func TileSetID(ds data.Store) (string, error) {
	mapboxConfig, err := ds.GetMapBoxConfig()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", mapBoxUsername, mapboxConfig.HeatMapTilesetName), nil
}
