package main


import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
	"log"
)

type MyEvent struct {
        Data []string `json:"data"`
}





func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	for _, element := range name.Data {
                str, ok := element.(string)
                if ok {
                        log.Printf("Data: %s", str)
                } else {
                        log.Printf("Received non-string data")
                }
        }
        return fmt.Sprintf("Hello %s!", name.Data ), nil
}

func main() {

	
	
	fmt.Sprintf("Hello!" )
        lambda.Start(HandleRequest)
}
