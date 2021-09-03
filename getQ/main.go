package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// func handler handles event
/* func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	return events.APIGatewayProxyResponse{Body: "Hello World! ", StatusCode: 200}, nil
} */

// HandleRequest is our handler func
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, _ = fmt.Println("parsed:", request.Body)
	return events.APIGatewayProxyResponse{Body: "response is working", StatusCode: 200}, nil
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

// func MapToApiGateway(w http.ResponseWriter, r *http.Request) (interface{}, error) {
// 	request := new(EmailResponderRequest)
// 	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
// 		return err.Error(), err
// 	}
// 	apiGatewayRequest := mapHttpRequestToGatewayRequest(*request)
// 	events, err := lambdahandler.HandleRequest(nil, apiGatewayRequest)
// 	if err != nil {
// 		return err.Error(), err
// 	}
// 	return events, nil
// }
