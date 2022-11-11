package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devilcove/httpclient"
)

type Response struct {
	IP string
}

func main() {
	var response Response
	var errResponse any
	endpoint := httpclient.JSONEndpoint[Response, any]{
		URL:           "https://api.ipify.org?format=json",
		Route:         "",
		Method:        http.MethodGet,
		Authorization: "",
		Data:          nil,
		Response:      response,
		ErrorResponse: errResponse,
	}
	answer, err := endpoint.GetJSON(response, errResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)

}
