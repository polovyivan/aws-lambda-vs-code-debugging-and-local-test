package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	awsRegion   string
	awsEndpoint string
	bucketName  string

	s3svc *s3.Client
)

func init() {
	awsRegion = os.Getenv("REGION")
	awsEndpoint = os.Getenv("BUCKET_ENDPOINT")
	bucketName = os.Getenv("BUCKET_NAME")

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	s3svc = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}

func uploadFileToS3(fileName string, fileContent string) {
    s3Key := fmt.Sprintf("%s.txt", fileName)   
    log.Printf("==>bucketName %q ",  bucketName)
    body := []byte(fileContent)
	resp, err := s3svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(s3Key),
		Body:               bytes.NewReader(body),
		ContentLength:      int64(len(body)),
		ContentType:        aws.String("application/text"),
		ContentDisposition: aws.String("attachment"),
	})

	log.Printf("S3 PutObject response: %+v", resp)

	if err != nil {
        log.Printf("Unable to upload %q to %q, %v", fileName, bucketName, err)
    } else {
        log.Printf("Successfully uploaded %q to %q\n", fileName, bucketName)
    }
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

    for _, message := range sqsEvent.Records {
        fmt.Println("Id", message.MessageId)
        fmt.Println("Source", message.EventSource)
        fmt.Println("Body", message.Body)
        uploadFileToS3(message.MessageId, message.Body);
    }

    return nil
}

func main() {
	lambda.Start(handler)
}