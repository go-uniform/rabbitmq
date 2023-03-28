package rabbitApi

import (
    "encoding/json"
    "fmt"
    "github.com/go-diary/diary"
    "net/http"
    "strings"
)

func (r *rabbitApi) QueueList() []Queue {
    uri := fmt.Sprintf("%s/queues", strings.TrimRight(r.BaseUri, "/"))

    client := &http.Client{}
    req, err := http.NewRequest("GET", uri, nil)
    req.SetBasicAuth(r.Username, r.Password)

    if err != nil {
        panic(err)
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")

    r.Page.Debug("rabbitApi.queue-list", diary.M{
        "method": "GET",
        "uri":    uri,
    })

    /* Handle Response */
    var queueListResponse []Queue

    /* Execute Request */
    body, statusCode, err := executeRequest(client, req)

    if statusCode != 200 {
        panic(fmt.Sprintf("api call '%s %s' failed with status '%d'", req.Method, uri, statusCode))
    }

    if err := json.Unmarshal(body, &queueListResponse); err != nil {
        panic(err)
    }

    return queueListResponse
}
