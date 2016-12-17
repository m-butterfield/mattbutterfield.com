package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/mattn/go-sqlite3"
)

const (
	awsRegion        = "us-east-1"
	bucketName       = "images.mattbutterfield.com"
	dbFileName       = "app.db"
	insertImageQuery = "INSERT INTO images (id) VALUES (?)"
	latestIDQuery    = "SELECT id FROM images ORDER BY id DESC LIMIT 1"
	maxKeys          = 100
)

var (
	db  *sql.DB
	svc *s3.S3
)

func main() {
	fmt.Println("Hello.")
	var err error
	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err)
	}
	svc = s3.New(session.New(&aws.Config{Region: aws.String(awsRegion)}))
	latestID, err := getLatestID()
	if err != nil {
		panic(err)
	}
	fmt.Println("Latest id: ", latestID)
	err = fetchImages(latestID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Program completed successfully!")
}

func getLatestID() (string, error) {
	fmt.Println("Fetching latest ID from database...")
	var id string
	err := db.QueryRow(latestIDQuery).Scan(&id)
	if err == sql.ErrNoRows {
		fmt.Println("No rows found in database, so fetching first key name from S3...")
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
	return id, err
}

func fetchImages(latestID string) error {
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
			err = storeKeyInDB(*result.Key)
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

func storeKeyInDB(keyName string) (err error) {
	fmt.Println("Storing info for key: ", keyName)
	_, err = db.Exec(insertImageQuery, keyName)
	return
}
