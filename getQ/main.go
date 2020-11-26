package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


func main() {
	lambda.Start(handler)
}

function handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	fmt.Println("Received body: ", request.Body)
return request.APIGatewayProxyResponse(Body:"Hello World! ", StatusCode: 200), nil

}