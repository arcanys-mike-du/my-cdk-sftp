package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type Event struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Role          string `json:"Role"`
	Policy        string `json:"Policy"`
	HomeDirectory string `json:"HomeDirectory"`
}

func HandleRequest(ctx context.Context, event Event) (Response, error) {
	// Replace these with your actual username and password logic
	const correctUsername = "testftpuser"
	const correctPassword = "test123456"

	// Fetch environment variables
	roleArn := os.Getenv("SFTP_ROLE_ARN")
	bucketName := os.Getenv("S3_BUCKET_NAME")
	homeDirectory := os.Getenv("SFTP_HOME_DIRECTORY")

	if event.Username == correctUsername && event.Password == correctPassword {
		return Response{
			Role:          roleArn,
			Policy:        fmt.Sprintf("{\"Version\": \"2012-10-17\",\"Statement\": [{\"Sid\": \"AllowFullAccessToBucket\",\"Action\": [\"s3:*\"],\"Effect\": \"Allow\",\"Resource\": [\"arn:aws:s3:::%s/*\"]}]}", bucketName),
			HomeDirectory: homeDirectory,
		}, nil
	}
	return Response{}, fmt.Errorf("authentication failed")
}

func main() {
	lambda.Start(HandleRequest)
}
