package main

"github.com/aws-aws-lambda-go/events"
"github.com/aws-aws-lambda-go/lambda"

func main() {
	lambda.Start(handler)
}

function handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){

return request.APIGatewayProxyResponse(Body:"Hello World! ", StatusCode: 200), nil

}