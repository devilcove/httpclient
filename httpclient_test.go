package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/Kairum-Labs/should"
)

func TestGetResponse(t *testing.T) {
	t.Run("ip endpoint", func(t *testing.T) {
		response, err := GetResponse(nil, http.MethodGet, "https://firefly.nusak.ca/ip", "", nil)
		should.BeNil(t, err)
		should.BeEqual(t, response.StatusCode, http.StatusOK)
		defer response.Body.Close()
		bytes, err := io.ReadAll(response.Body)
		should.BeNil(t, err)
		ip := net.ParseIP(string(bytes))
		if ip == nil {
			t.Log(string(bytes))
			t.Fail()
		}
	})
	t.Run("invalid endpoint", func(t *testing.T) {
		response, err := GetResponse(nil, http.MethodGet, "https://firefly.nusak.ca/invalidendpoint", "", nil)
		should.BeNil(t, err)
		should.BeEqual(t, response.StatusCode, http.StatusNotFound)
	})
	t.Run("login and hello", func(t *testing.T) {
		type Data struct {
			User string
			Pass string
		}
		var data Data
		data.User = "demo"
		data.Pass = "pass"
		jwt := struct {
			JWT string
		}{}
		response, err := GetResponse(data, http.MethodPost, "https://firefly.nusak.ca/login", "", nil)
		should.BeNil(t, err)
		defer response.Body.Close()
		should.BeNil(t, json.NewDecoder(response.Body).Decode(&jwt))
		response, err = GetResponse("", http.MethodGet, "https://firefly.nusak.ca/api/hello", jwt.JWT, nil)
		should.BeNil(t, err)
		should.BeEqual(t, response.StatusCode, http.StatusOK)
		defer response.Body.Close()
		bytes, err := io.ReadAll(response.Body)
		should.BeNil(t, err)
		ip := net.ParseIP(string(bytes))
		if ip == nil {
			t.Fail()
		}
	})

	t.Run("badlogin", func(t *testing.T) {
		data := struct {
			User string
			Pass string
		}{
			User: "demo",
			Pass: "badpass",
		}
		answer := struct {
			Message string
		}{}
		response, err := GetResponse(data, http.MethodPost, "https://firefly.nusak.ca/login", "", nil)
		should.BeNil(t, err)
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		should.BeNil(t, err)
		should.BeEqual(t, response.StatusCode, http.StatusBadRequest)
		should.BeNil(t, json.Unmarshal(body, &answer))
		should.BeEqual(t, answer.Message, "invalid username or password")
	})
	t.Run("hello", func(t *testing.T) {
	})
}

func TestGetJSON(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		data := struct {
			User string
			Pass string
		}{
			User: "demo",
			Pass: "pass",
		}
		response := struct {
			JWT string
		}{}
		var errResponse any
		e := JSONEndpoint[struct{ JWT string }, any]{
			URL:           "http://firefly.nusak.ca",
			Route:         "/login",
			Method:        http.MethodPost,
			Data:          data,
			Response:      response,
			ErrorResponse: errResponse,
		}
		answer, errs, err := e.GetJSON(response, errResponse)
		should.BeNil(t, err)
		answerType := fmt.Sprintf("%T", answer)
		t.Log(answerType)
		should.BeTrue(t, answerType == "struct { JWT string }")
		should.BeNil(t, errs)
	})
}

type dummyStruct struct {
	Value string `json:"value"`
}

func TestGetJSON_PanicOnNilResponse(t *testing.T) {
	var resp dummyStruct
	var errResp dummyStruct

	_, _, err := GetJSON(nil, resp, errResp, "GET", ":", "", nil)

	if err == nil {
		t.Fatal("err should not be nil")
	}

}
