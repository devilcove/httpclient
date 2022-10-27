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
	endpoint := httpclient.JSONEndpoint[Response]{
		URL:           "https://api.ipify.org?format=json",
		Route:         "",
		Method:        http.MethodGet,
		Authorization: "",
		Data:          nil,
		Response:      response,
	}
	answer, code, err := endpoint.GetJSON(response)
	if err != nil {
		log.Fatal(err)
	}
	if code != http.StatusOK {
		log.Fatal(err, code)
	}
	fmt.Println(answer)

}
