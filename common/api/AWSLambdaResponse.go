package api

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func GetAWSResponse(response string, err error) events.APIGatewayProxyResponse {
	var responseBody string
	var statusCode int

	if err != nil {
		responseBody, statusCode = err.Error(), 500
	} else {
		responseBody, statusCode = response, 200
	}

	ApiResponse := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": os.Getenv("CORS"),
		},
		Body:       responseBody,
		StatusCode: statusCode,
	}
	return ApiResponse
}
