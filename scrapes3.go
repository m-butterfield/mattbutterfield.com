package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/m-butterfield/mattbutterfield.com/datastore"
	"github.com/m-butterfield/mattbutterfield.com/website"
)

const (
	awsRegion  = "us-east-1"
	bucketName = "images.mattbutterfield.com"
	maxKeys    = 100
)

func main() {
	db, err := datastore.InitDB(website.DBFileName)
	if err != nil {
		panic(err)
	}
	store := datastore.DBImageStore{DB: db}
	svc := s3.New(session.New(&aws.Config{Region: aws.String(awsRegion)}))
	latestID, err := getLatestID(store, svc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Latest id: ", latestID)
	err = fetchImages(store, svc, latestID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Program completed successfully!")
}

func getLatestID(store datastore.ImageStore, svc *s3.S3) (string, error) {
	image, err := store.GetLatestImage()
	if err == sql.ErrNoRows {
		result, err := svc.ListObjects(&s3.ListObjectsInput{
			Bucket:  aws.String(bucketName),
			MaxKeys: aws.Int64(1),
		})
		if err != nil {
			return "", err
		}
		if len(result.Contents) == 0 {
			return "", errors.New("No keys found in bucket.")
		}
		return *result.Contents[0].Key, nil
	}
	if err != nil {
		return "", err
	}
	return image.ID, err
}

func fetchImages(store datastore.ImageStore, svc *s3.S3, latestID string) error {
	fmt.Println("Fetching new keys from S3...")
	firstRunThrough := true
	for {
		result, err := svc.ListObjects(&s3.ListObjectsInput{
			Bucket:  aws.String(bucketName),
			Marker:  aws.String(latestID),
			MaxKeys: aws.Int64(maxKeys),
		})
		if err != nil {
			return err
		}
		if len(result.Contents) == 0 {
			if firstRunThrough {
				fmt.Println("No new keys found. You're all up to date!")
			}
			break
		}
		for _, result := range result.Contents {
			fmt.Println("Saving image: ", *result.Key)
			err = store.SaveImage(datastore.Image{ID: *result.Key, Caption: ""})
			if err != nil {
				return err
			}
			latestID = *result.Key
		}
		firstRunThrough = false
	}
	fmt.Println("Done fetching images.")
	return nil
}
