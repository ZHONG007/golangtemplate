package main


import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
	"log"
	"io/ioutil"
	"net/http"
)

type MyEvent struct {
        Data string `json:"data"`
}




func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("Received message: %s", name.Data)
        return fmt.Sprintf("Hello %s!", name.Data ), nil
}

func main() {

	
	
	fmt.Sprintf("Hello!" )
        lambda.Start(HandleRequest)
}
