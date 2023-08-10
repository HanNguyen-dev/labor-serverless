package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	BLS_URL         = os.Getenv("BLS_API")
	TIME_SERIES_URL = BLS_URL + "/timeseries/data"
)

func GetCPI() (string, error) {
	URL := TIME_SERIES_URL + "/CUUR0000SA0"
	return BLSApiHandler(URL, "getInflation")
}

func GetUnemployment() (string, error) {
	URL := TIME_SERIES_URL + "/LNS14000000"
	return BLSApiHandler(URL, "getUnemployment")
}

func BLSApiHandler(URL string, requestName string) (string, error) {
	resp, err := http.Get(URL)

	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error in making the %v request", requestName))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error in reading the response from %v", requestName))
	}
	return string(body), nil
}
