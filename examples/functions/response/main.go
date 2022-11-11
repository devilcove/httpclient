package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/devilcove/httpclient"
)

func main() {
	response, err := httpclient.GetResponse(nil, http.MethodGet, "http://ifconfig.me", "", nil)
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
