package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	request := events.APIGatewayProxyRequest{}

	expectedResponse := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "CORS",
		},
		Body:       "responseBody",
		StatusCode: 200,
	}
	GetAWSResponse = func(request string, err error) events.APIGatewayProxyResponse {
		return expectedResponse
	}
	actualResponse, err := HandleRequest(request)
	if err != nil {
		t.Errorf("HandleRequest errored out")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("Expected Response doesn't equal Actual") //, actualResponse)
	}
}
