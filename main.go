package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.KinesisEvent) error {
	fmt.Printf("Data: %s\n", event.Records[0].Kinesis.Data)

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
