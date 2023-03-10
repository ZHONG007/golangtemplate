package main


import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Message string `json:"message"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
        return fmt.Sprintf("Hello %s!", name.Message ), nil
}

func main() {
	fmt.Sprintf("Hello!" )
        lambda.Start(HandleRequest)
}
