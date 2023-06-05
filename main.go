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
        Message string `json:"message"`
}




func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("Received message: %s", name.Message)
        response, err := http.Get("https://8080-cs-476722282134-default.cs-europe-west1-iuzs.cloudshell.dev/sensors")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
        return fmt.Sprintf("Hello %s!", name.Message ), nil
}

func main() {

	
	
	fmt.Sprintf("Hello!" )
        lambda.Start(HandleRequest)
}
