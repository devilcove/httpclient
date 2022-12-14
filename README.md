# httpclient 
* simple library to facilitate calls to http endpoints
## two methods to call
* function with parameters
* method with structure
## two options for results
* raw response (*http.Response)
* response body decoded to json 
# Examples
see examples directory
## Full Response
equivalent of 

`curl 'https://api.example.com' -H 'Authorization: Bearer API_TOKEN'`
### function 
```
func main() {
    token := os.Getenv("API_TOKEN")
    response, err := httpclient.GetResponse("", http.MethodGet, "https://api.example.com", token)
    if err !=nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    bytes, err := io.ReadAll(response.Body)
    log.Println(response.StatusCode, string(bytes))
}
```

### method 
```
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
```
## JSON response
equivalent of 

`curl 'https://api.example.com/login' -H 'Authorization: Bearer API_TOKEN' - H 'Content-Type: application/json' -d '{"name":"admin","pass":"password"}'`
### function 
```
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
    response, err := httpclient.GetJSON(data, http.MethodPOST, "https://api.example.com/api/login", "Bearer " + token)
    if err !=nil {
        log.Fatal(err)
    }
    log.Println(response)
}
```
### method
```
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
        URL: "https://api.example.com",
	    Method: http.MethodPOST
	    Authorization: token,
        Route: "/api/login",
        Data: data,
        Response: response,
    }
    jsonEndpoint := httpclient.JSONEndpoint[struct {Language_Pref string JWT string}] {
        endpoint,
        response,
    }
    answer, err := httpclient.JSON(response)
    if err !=nil {
        log.Fatal(err)
    }
    log.Println(response)
    ```