package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getInflation() (string, error) {
	URL := os.Getenv("BLS_API") + "/timeseries/data/CUUR0000SA0"

	resp, err := http.Get(URL)

	if err != nil {
		return "", fmt.Errorf("Error in fetching inflation data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("Error in reading the inflation data")
	}
	return string(body), nil
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	blsResponse, err := getInflation()

	var response string
	var statusCode int
	if err != nil {
		response = err.Error()
		statusCode = 500
	} else {
		response = blsResponse
		statusCode = 200
	}

	ApiResponse := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": os.Getenv("CORS"),
		},
		Body:       response,
		StatusCode: statusCode,
	}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
