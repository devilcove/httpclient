package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devilcove/httpclient"
)

type Answer struct {
	IP string
}

// equivalient of curl api.ipify.org?format=json
func main() {
	var answer Answer
	ip, err := httpclient.GetJSON(nil, answer, http.MethodGet, "http://api.ipify.org?format=json", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)
}
