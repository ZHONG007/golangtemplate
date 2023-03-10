package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, payload []byte) {
	fmt.Printf("Received message: %s\n", string(payload))
}

func main() {
	lambda.Start(handler)
}
