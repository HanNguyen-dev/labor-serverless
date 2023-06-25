package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hannguyen-dev/labor-serverless/common/api"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	blsResponse, err := api.GetUnemployment()
	return api.GetAWSResponse(blsResponse, err), nil
}

func main() {
	lambda.Start(HandleRequest)
}
