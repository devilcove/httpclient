package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/devilcove/httpclient"
)

type Answer struct {
	IP string
}

// equivalient of curl api.ipify.org?format=json
func main() {
	var answer Answer
	var errResp any
	ip, _, err := httpclient.GetJSON(nil, answer, errResp, http.MethodGet, "http://api.ipify.org?format=json", "", nil)
	if err != nil {
		if strings.Contains(err.Error(), "non ok status") {
			fmt.Println(ip, err)
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(ip)
}
