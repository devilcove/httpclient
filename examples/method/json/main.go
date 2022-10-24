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
	end := httpclient.Endpoint{
		URL:           "https://api.ipify.org?format=json",
		Route:         "",
		Method:        http.MethodGet,
		Authorization: "",
		Data:          nil,
	}
	endpoint := httpclient.JSONEndpoint[Response]{
		Endpoint: end,
		Response: response,
	}
	answer, err := endpoint.GetJSON(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)

}
