package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hannguyen-dev/labor-serverless/common/api"
	"github.com/hannguyen-dev/labor-serverless/common/model"
)

func getInflation() (string, error) {
	response, err := api.GetCPI()
	var parsedBody model.BLSResponse

	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(response), &parsedBody)

	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error in parsing JSON into object"))
	}

	// Calculate inflation rates
	if parsedBody.Status != "REQUEST_NOT_PROCESSED" {
		seriesData := parsedBody.Results.Series[0].Data

		inflationRates := []model.InflationRates{}

		for i := len(seriesData) - 1; i > 11; i-- {
			lastYearData := seriesData[i]
			currentYearData := seriesData[i-12]

			lastYearCPI, errLastYear := strconv.ParseFloat(lastYearData.Value, 32)
			currentYearCPI, errCurrentYear := strconv.ParseFloat(currentYearData.Value, 32)

			if errLastYear != nil || errCurrentYear != nil {
				return "", fmt.Errorf(fmt.Sprint("Error in converting CPI string into float"))
			}

			rate := 100 * (currentYearCPI - lastYearCPI) / lastYearCPI
			dataPoint := model.InflationRates{
				X: fmt.Sprintf("%s %s", currentYearData.PeriodName[:3], currentYearData.Year),
				Y: math.Round(rate*100) / 100,
			}

			inflationRates = append(inflationRates, dataPoint)
		}

		stringifiedData, err := json.Marshal(inflationRates)
		if err != nil {
			return "", fmt.Errorf(fmt.Sprint("Error in stringifying inflation rates"))
		}
		return string(stringifiedData), nil
	}
	return "", fmt.Errorf(fmt.Sprint("BLS didn't process request"))
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	blsResponse, err := getInflation()
	return api.GetAWSResponse(blsResponse, err), nil
}

func main() {
	lambda.Start(HandleRequest)
}
