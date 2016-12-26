package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/m-butterfield/mattbutterfield.com/datastore"
)

const (
	awsRegion  = "us-east-1"
	bucketName = "images.mattbutterfield.com"
	maxKeys    = 100
)

func main() {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(awsRegion)}))
	latestID, err := getLatestID(svc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Latest id: ", latestID)
	err = fetchImages(svc, latestID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Program completed successfully!")
}

func getLatestID(svc *s3.S3) (string, error) {
	latestID, err := datastore.GetLatestID()
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
	return latestID, err
}

func fetchImages(svc *s3.S3, latestID string) error {
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
			err = datastore.SaveImage(*result.Key, nil)
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
