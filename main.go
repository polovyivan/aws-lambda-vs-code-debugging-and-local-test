package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

    for _, message := range sqsEvent.Records {
        fmt.Println("Id", message.MessageId)
        fmt.Println("Source", message.EventSource)
        fmt.Println("Body", message.Body)
    }

    return nil
}

func main() {
    lambda.Start(handler)
}