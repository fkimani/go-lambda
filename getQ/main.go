package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest is our handler func
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, _ = fmt.Println("parsed:", request.Body)
	return events.APIGatewayProxyResponse{Body: "hello, response is working", StatusCode: 200}, nil
}

func main() {
	environment := os.Getenv("ENV")
	if environment == "dev" {
		// router.NewRouter()
		select {}
	} else {
		lambda.Start(HandleRequest)
	}
}
