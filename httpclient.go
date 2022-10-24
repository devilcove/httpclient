package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kr/pretty"
)

var Client http.Client

// Endpoint struct for an http endpoint
type Endpoint struct {
	URL           string
	Method        string
	Route         string
	Authorization string
	Data          any
}

type JSONEndpoint[T comparable] struct {
	Endpoint
	Response T
}

func init() {
	Client = http.Client{
		Timeout: 30 * time.Second,
	}
}

// GetResponse returns respnse from http request to url
func GetResponse(data any, method, url, auth string) (*http.Response, error) {
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
		request.Header.Set("Authorization", auth)
	}
	return Client.Do(request)
}

// JSON returns JSON response from http request
func GetJSON[T any](data any, resp T, method, url, auth string) (any, error) {
	response, err := GetResponse(data, method, url, auth)
	if err != nil {
		return nil, err
	}
	pretty.Println("response before json decodiing", resp)
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}
	pretty.Println("response after json decodiing", resp)
	return resp, nil
}

// GetResponse returns response from http endpoint
func (e *Endpoint) GetResponse() (*http.Response, error) {
	return GetResponse(e.Data, e.Method, e.URL+e.Route, e.Authorization)
}

// GetJSON returns JSON received from http endpoint
func (e *JSONEndpoint[T]) GetJSON(response T) (any, error) {
	return GetJSON(e.Data, response, e.Method, e.URL+e.Route, e.Authorization)
}
