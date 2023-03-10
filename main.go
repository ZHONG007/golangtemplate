package main


import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.IoTEvent) {
	fmt.Printf("Received mqtt message: %s\n", event.Payload)
}

func main() {
	lambda.Start(handler)
}
