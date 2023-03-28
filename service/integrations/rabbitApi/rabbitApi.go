package rabbitApi

import (
    "github.com/go-diary/diary"
)

type rabbitApi struct {
    Page     diary.IPage
    BaseUri  string
    Username string
    Password string
}

type IRabbitApi interface {
    QueueList() []Queue
}

func NewRabbitApiConnector(page diary.IPage, uri, username, password string) IRabbitApi {
    var instance IRabbitApi
    page.Scope("infobip", func(p diary.IPage) {
        instance = &rabbitApi{
            Page:     page,
            BaseUri:  uri,
            Username: username,
            Password: password,
        }
    })
    return instance
}
