httpclient is a simple library to facilitate calls to http endpoints.
two methods to call
    function with parameters
    method with structure
two options for results
    raw response (*http.Response)
    response body decoded to json

example function call with parameters (full response)
equivalent of curl 'https://api.example.com' -H 'Authorization: Bearer API_TOKEN'


func main() {
    token := os.Getenv("API_TOKEN")
    response, err := httpclient.API("", http.MethodGet, "https://api.example.com", token)
    if err !=nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    bytes, err := io.ReadAll(response.Body)
    log.Println(response.StatusCode, string(bytes))
}

example method call with parameters (full response)
equivalent of curl 'https://api.example.com' -H 'Authorization: Bearer API_TOKEN'

func main() {
    token := os.Getenv("API_TOKEN")
    endpoint := httpclient.Endpoint {
        URL: "https://api.clustercat.com",
	    Method: http.MethodGet    
	    Authorization: token,
    }
    response, err := endpoint.GetResponse()
    defer response.Body.Close()
    bytes, err := io.ReadAll(response.Body)
    log.Println(response.StatusCode, string(bytes))
}

example function call with json response
equivalent of curl 'https://api.example.com/api/login' -H 'Authorization: Bearer API_TOKEN' - H 'Content-Type: application/json' -d '{"name":"admin","pass":"password"}'

func main() {
    token := os.Getenv("API_TOKEN")
    data := struct {
        Name: "admin",
        Pass: "password",
    }
    response := struct {
        Language_Pref: "EN",
        JWT: "some string",
    }
    response, err := httpclient.JSON(data, http.MethodPOST, "https://api.example.com/api/login", token)
    if err !=nil {
        log.Fatal(err)
    }
    log.Println(response)
}

example method call with json response
equivalent of curl 'https://api.example.com/api/login' -H 'Authorization: Bearer API_TOKEN' - H 'Content-Type: application/json' -d '{"name":"admin","pass":"password"}'

func main() {
    token := os.Getenv("API_TOKEN")
    data := struct {
        Name: "admin",
        Pass: "password",
    }
    response := struct {
        Language_Pref: "EN",
        JWT: "some string",
    }
    endpoint := httpclient.Endpoint {
        URL: "https://api.clustercat.com",
	    Method: http.MethodPOST
	    Authorization: token,
        Route: "/api/login",
        Data: data,
        Response: response,
    }
    response, err := httpclient.JSON()
    if err !=nil {
        log.Fatal(err)
    }
    log.Println(response)
}

