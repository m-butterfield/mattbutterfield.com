package main

import (
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/mattn/go-sqlite3"
)

const (
	awsRegion  = "us-east-1"
	bucketName = "images.mattbutterfield.com"
	dbFileName = "app.db"
	maxKeys    = 100
)

var (
	db *sql.DB
)

func main() {
	fmt.Println("Hello.")
	var err error
	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err)
	}
	latestID, err := getLatestID()
	if err != nil {
		panic(err)
	}
	fmt.Println("Latest id: ", latestID)
	err = fetchImages(latestID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}

func getLatestID() (id string, err error) {
	err = db.QueryRow("SELECT id FROM images ORDER BY created_at DESC LIMIT 1").Scan(&id)
	return
}

func fetchImages(latestID string) error {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(awsRegion)}))
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
			break
		}
		fmt.Println("Keys:")
		for _, result := range result.Contents {
			fmt.Printf("%s : %s\n", aws.StringValue(result.Key), result.LastModified)
			err = storeKeyInDB(*result.Key)
			if err != nil {
				return err
			}
			latestID = *result.Key
		}
	}
	return nil
}

func storeKeyInDB(keyName string) error {
	return nil
}
