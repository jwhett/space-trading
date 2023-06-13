package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiBaseUrl = "https://api.spacetraders.io/v2"

var apiToken = os.Getenv("ST_TOKEN")

func makeRequest(method, url string, reader io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
