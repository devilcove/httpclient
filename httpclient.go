package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var Client http.Client

func init() {
	Client = http.Client{
		Timeout: 30 * time.Second,
	}
}

func API(data, resp any, method, url, auth string) (any, error) {
	var request *http.Request
	var err error
	if data != "" {
		payload, err := json.Marshal(data)
		if err != nil {
			return request, fmt.Errorf("error encoding data %w", err)
		}
		request, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
		if err != nil {
			return request, fmt.Errorf("error creating http request %w", err)
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return request, fmt.Errorf("error creating http request %w", err)
		}
	}
	if auth != "" {
		request.Header.Set("Authorization", "Bearer "+auth)
	}
	response, err := Client.Do(request)
	if err != nil {
		return response, err
	}
	defer response.Body.Close()
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return data, fmt.Errorf("endpoint return status %s", response.Status)
	}
	if err = json.Unmarshal(data.([]byte), resp); err != nil {
		return nil, fmt.Errorf("response not json %w", err)
	}
	return resp, nil
}
