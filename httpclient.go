package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var Client http.Client
var ErrJSON = errors.New("httpclient: json error")
var ErrRequest = errors.New("httpclient: error creating http request")
var ErrStatus = errors.New("httpclient: http.Status not OK")

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
			return response, ErrJSON
		}
		request, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
		if err != nil {
			return response, ErrRequest
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return response, ErrRequest
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
// if response.code is http.StatusOK (200) returns body decoded to resp
// if response.code is not http.Status ok returns response(*http.Response) with err set to
// 'non ok response code'
// if json decode err returns response and err set to 'json decode err'
// for any other err returns nil and err
func GetJSON[T any, R any](data any, resp T, errResponse R, method, url, auth string, headers []Header) (any, error) {
	response, err := GetResponse(data, method, url, auth, headers)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		if err := json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			return response, ErrJSON
		}
		return errResponse, ErrStatus
	}
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return response, ErrJSON
	}
	return resp, nil
}

// GetResponse returns response from http endpoint
func (e *Endpoint) GetResponse() (*http.Response, error) {
	return GetResponse(e.Data, e.Method, e.URL+e.Route, e.Authorization, e.Headers)
}

// GetJSON returns JSON received from http endpoint
func (e *JSONEndpoint[T, R]) GetJSON(response T, errResponse R) (any, error) {
	return GetJSON(e.Data, response, errResponse, e.Method, e.URL+e.Route, e.Authorization, e.Headers)
}
