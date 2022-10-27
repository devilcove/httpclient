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
	ip, code, err := httpclient.GetJSON(nil, answer, http.MethodGet, "http://api.ipify.org?format=json", "")
	if err != nil {
		log.Fatal(err)
	}
	if code != http.StatusOK {
		log.Fatal(err, code)
	}
	fmt.Println(ip)
}
