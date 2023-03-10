package main
import (
        "context"
        "fmt"
        "net/http"
    
        "github.com/aws/aws-lambda-go/events"
        "github.com/aws/aws-lambda-go/lambda"
    )
 
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        fmt.Println("Received request:", request)
    
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       "hello from golang in aws lambda",
        }, nil
    }
 
    func main() {
        fmt.Println("Received request start")
        lambda.Start(handler)
    }