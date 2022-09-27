package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var Client http.Client

func init() {
	Client = http.Client{
		Timeout: 30 * time.Second,
	}
}

func API(data, resp any, method, url, auth string) (*http.Response, error) {
	var request *http.Request
	var response *http.Response
	var err error
	if data != "" {
		payload, err := json.Marshal(data)
		if err != nil {
			return response, fmt.Errorf("error encoding data %w", err)
		}
		request, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
		if err != nil {
			return response, fmt.Errorf("error creating http request %w", err)
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return response, fmt.Errorf("error creating http request %w", err)
		}
	}
	if auth != "" {
		request.Header.Set("Authorization", "Bearer "+auth)
	}
	return Client.Do(request)
}
