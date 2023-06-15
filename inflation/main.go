package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getInflation() string {
	URL := os.Getenv("BLS_API") + "/timeseries/data/CUUR0000SA0"

	resp, err := http.Get(URL)

	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return string(body)
		}
	}
	return ""
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := getInflation()
	ApiResponse := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": os.Getenv("CORS"),
		},
		Body:       response,
		StatusCode: 200,
	}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
