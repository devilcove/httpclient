package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Client http.Client

const (
	ErrJSON    = "httpclient: json error decoding response"
	ErrJSON2   = "httpclient: json error decoding error response"
	ErrRequest = "httpclient: error creating http request"
	ErrStatus  = "httpclient: http.Status not OK"
)

// ErrHttpClient error stuct
type ErrHttpClient struct {
	Status  int
	Message string
	Body    http.Response
}

// ErrHttpClient.Error implementation of error interface
func (e ErrHttpClient) Error() string {
	return e.Message
}

// Endpoint struct for an http endpoint
type Endpoint struct {
	URL           string
	Method        string
	Route         string
	Authorization string
	Headers       []Header
	Data          any
}

type JSONEndpoint[T any, R any] struct {
	URL           string
	Method        string
	Route         string
	Authorization string
	Headers       []Header
	Data          any
	Response      T
	ErrorResponse R
}

type Header struct {
	Name  string
	Value string
}

func init() {
	Client = http.Client{
		Timeout: 30 * time.Second,
	}
}

// GetResponse returns respnse from http request to url
func GetResponse(data any, method, url, auth string, headers []Header) (*http.Response, error) {
	var request *http.Request
	var response *http.Response
	var err error
	if data != nil {
		payload, err := json.Marshal(data)
		if err != nil {
			return response, ErrHttpClient{
				Status:  http.StatusBadRequest,
				Message: ErrJSON,
				Body:    *response,
			}
		}
		request, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
		if err != nil {
			log.Println("httclient:NewRequest with payload", err)
			return response, ErrHttpClient{
				Status:  http.StatusBadRequest,
				Message: ErrRequest,
				Body:    *response,
			}
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			log.Println("httclient:NewRequest", err)
			return response, ErrHttpClient{
				Status:  http.StatusBadRequest,
				Message: ErrRequest,
				Body:    *response,
			}
		}
	}
	if auth != "" {
		request.Header.Set("Authorization", auth)
	}
	for _, header := range headers {
		request.Header.Set(header.Name, header.Value)
	}
	return Client.Do(request)
}

// JSON returns JSON response from http request
// if response.code is http.StatusOK (200) returns resp, nil, nil
// if response.code is not http.Status ok returns nil, errResponse, ErrStatus
// if json decode err returns  nil, nil, err set to 'json decode err'
// for any other err returns nil, nil, err
func GetJSON[T any, R any](data any, resp T, errResponse R, method, url, auth string, headers []Header) (T, R, error) {
	response, err := GetResponse(data, method, url, auth, headers)
	if err != nil {
		return resp, errResponse, ErrHttpClient{
			Status:  response.StatusCode,
			Message: fmt.Errorf("%w", err).Error(),
			Body:    *response,
		}
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		if err := json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			return resp, errResponse, ErrHttpClient{
				Status:  response.StatusCode,
				Message: fmt.Errorf("json decode err with error response %w", err).Error(),
				Body:    *response,
			}
		}
		return resp, errResponse, ErrHttpClient{
			Status:  response.StatusCode,
			Message: ErrStatus,
			Body:    *response,
		}
	}
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return resp, errResponse, ErrHttpClient{
			Status:  response.StatusCode,
			Message: fmt.Errorf("json decode err with response %w", err).Error(),
			Body:    *response,
		}
	}
	return resp, errResponse, nil
}

// GetResponse returns response from http endpoint
func (e *Endpoint) GetResponse() (*http.Response, error) {
	return GetResponse(e.Data, e.Method, e.URL+e.Route, e.Authorization, e.Headers)
}

// GetJSON returns JSON received from http endpoint
func (e *JSONEndpoint[T, R]) GetJSON(response T, errResponse R) (T, R, error) {
	return GetJSON(e.Data, response, errResponse, e.Method, e.URL+e.Route, e.Authorization, e.Headers)
}
