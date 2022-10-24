package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/devilcove/httpclient"
)

func main() {
	endpoint := httpclient.Endpoint{
		URL:           "http://ifconfig.me",
		Route:         "",
		Method:        http.MethodGet,
		Authorization: "",
		Data:          nil,
	}
	response, err := endpoint.GetResponse()
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %v your ip: %v", response.Status, string(bytes))
}
