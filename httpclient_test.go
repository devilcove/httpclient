package httpclient

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/matryer/is"
)

func TestAPI(t *testing.T) {
	t.Run("ip endpoint", func(t *testing.T) {
		response, err := API(nil, http.MethodGet, "https://firefly.nusak.ca/ip", "")
		is := is.New(t)
		is.NoErr(err)
		is.Equal(response.StatusCode, http.StatusOK)
		defer response.Body.Close()
		bytes, err := io.ReadAll(response.Body)
		is.NoErr(err)
		ip := net.ParseIP(string(bytes))
		if ip == nil {
			is.Fail()
		}
	})
	t.Run("invalid endpoint", func(t *testing.T) {
		response, err := API(nil, http.MethodGet, "https://firefly.nusak.ca/invalidendpoint", "")
		is := is.New(t)
		is.NoErr(err)
		is.Equal(response.StatusCode, http.StatusNotFound)
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
		response, err := API(data, http.MethodPost, "https://firefly.nusak.ca/login", "")
		is := is.New(t)
		is.NoErr(err)
		defer response.Body.Close()
		is.NoErr(json.NewDecoder(response.Body).Decode(&jwt))
		response, err = API("", http.MethodGet, "https://firefly.nusak.ca/api/hello", jwt.JWT)
		is.NoErr(err)
		is.Equal(response.StatusCode, http.StatusOK)
		defer response.Body.Close()
		bytes, err := io.ReadAll(response.Body)
		is.NoErr(err)
		ip := net.ParseIP(string(bytes))
		if ip == nil {
			is.Fail()
		}
	})

	t.Run("badlogin", func(t *testing.T) {
		type Data struct {
			User string
			Pass string
		}
		answer := struct {
			Request Data
			Error   string
		}{}
		var data Data
		data.Pass = "badpass"
		response, err := API(data, http.MethodPost, "https://firefly.nusak.ca/login", "")
		is := is.New(t)
		is.NoErr(err)
		defer response.Body.Close()
		is.Equal(response.StatusCode, http.StatusBadRequest)
		is.NoErr(json.NewDecoder(response.Body).Decode(&answer))
		is.Equal(answer.Error, "invalid username or password")
	})
	t.Run("hello", func(t *testing.T) {

	})

}
