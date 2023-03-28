package rabbitApi

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

// wrapper function to execute external requests (used for mocking in unit testing)
var executeRequest = func(client *http.Client, req *http.Request) ([]byte, int, error) {
    res, err := client.Do(req)
    if err != nil {
        return nil, -1, err
    }

    var body []byte = nil
    if res.Body != nil {
        defer res.Body.Close()
        body, err = ioutil.ReadAll(res.Body)
        if err != nil {
            panic(err)
        }
    }

    return body, res.StatusCode, err
}

// wrapper function to print external request (used for virtual mode used for local environment runs)
var printRequest = func(client *http.Client, req *http.Request) {
    var bodyData []byte
    if req.Body != nil {
        data, err := ioutil.ReadAll(req.Body)
        if err != nil {
            panic(err)
        }
        bodyData = data
    }
    fmt.Println()
    fmt.Println()
    fmt.Println()
    fmt.Println("########################################################################################################################################################################################")
    fmt.Printf("%s %s\n", req.Method, req.URL)
    fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
    var isJsonBody = false
    for name, values := range req.Header {
        for _, value := range values {
            if strings.ToLower(name) == "content-type" && strings.ToLower(value) == "application/json" {
                isJsonBody = true
            }
            fmt.Printf("%s: %s\n", name, value)
        }
    }
    if len(bodyData) > 0 {
        fmt.Println("****************************************************************************************************************************************************************************************")
        var printed = false
        if isJsonBody {
            var temp map[string]interface{}
            err := json.Unmarshal(bodyData, &temp)
            if err == nil {
                out, err := json.MarshalIndent(temp, "", "    ")
                if err == nil && out != nil {
                    fmt.Print(string(out))
                    printed = true
                }
            }
        }
        if !printed {
            fmt.Println(string(bodyData))
        }
    }
    fmt.Println("########################################################################################################################################################################################")
    fmt.Println()
    fmt.Println()
    fmt.Println()
}
