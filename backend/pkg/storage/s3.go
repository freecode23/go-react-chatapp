package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var region string = "us-east-2"
var bucketName string = "gochat-bucket"

func getCurrentDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02") // YYYY-MM-DD format
}

func getCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("15:04:05") // HH-mm-ss format
}

func SaveChatHistory(chatHistoryStr []string) error {

	fmt.Println("save chat hist:", chatHistoryStr)
	// Convert the slice of strings to a single JSON string
	chatHistoryJSONBytes, err := json.Marshal(chatHistoryStr)
	if err != nil {
		return err
	}

	// 2. init session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)

	// 3. create key for the data
	datestamp := getCurrentDate()
	timestamp := getCurrentTime()
	keyStr := fmt.Sprintf("chat-history/%s/%s.json", datestamp, timestamp)

	// 4. insert to s3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyStr), // or generate a dynamic name if needed
		Body:   bytes.NewReader(chatHistoryJSONBytes),
	})

	return err
}
