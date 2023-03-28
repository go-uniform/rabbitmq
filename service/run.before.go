package service

import (
    "github.com/go-diary/diary"
    "service/service/info"
    "service/service/integrations/rabbitAmqp"
    "service/service/integrations/rabbitApi"
    "sync"
)

func RunBefore(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
    if info.Virtualize {
        panic("this service can't be started in virtual mode, it is not yet supported")
    }

    amqpUri, ok := info.Args["amqpUri"].(string)
    if !ok {
        panic("amqpUri must be a string")
    }
    amqpsCert, ok := info.Args["amqpsCert"].(string)
    if !ok {
        panic("amqpsCert must be a string")
    }
    amqpsKey, ok := info.Args["amqpsKey"].(string)
    if !ok {
        panic("amqpsKey must be a string")
    }
    disableAmqps, ok := info.Args["disableAmqps"].(bool)
    if !ok {
        disableAmqps = false
    }
    apiUri, ok := info.Args["apiUri"].(string)
    if !ok {
        panic("apiUri must be a string")
    }
    apiUsername, ok := info.Args["apiUsername"].(string)
    if !ok {
        panic("apiUsername must be a string")
    }
    apiPassword, ok := info.Args["apiPassword"].(string)
    if !ok {
        panic("apiPassword must be a string")
    }
    workersPerQueue, ok := info.Args["workersPerQueue"].(int)
    if !ok {
        panic("workersPerQueue must be a string")
    }

    info.WorkersPerQueue = workersPerQueue
    info.RabbitApi = rabbitApi.NewRabbitApiConnector(p, apiUri, apiUsername, apiPassword)

    uri, config := rabbitAmqp.CompileRabbitAmqpConfig(amqpUri, amqpsCert, amqpsKey, disableAmqps)
    info.RabbitAmqp = rabbitAmqp.NewRabbitAmqpConnector(p, uri, config)
    info.Shutdown = shutdown
    info.Group = group
    info.DiaryPage = p
}
