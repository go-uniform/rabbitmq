package service

import (
    "github.com/go-diary/diary"
    rabbit "github.com/wagslane/go-rabbitmq"
    "service/service/events"
    "service/service/info"
    "service/service/integrations/rabbitApi"
    "service/service/utils"
    "sync"
)

func RunAfter(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
    var queues = info.RabbitApi.QueueList()
    if queues == nil {
        queues = make([]rabbitApi.Queue, 0)
    }
    queues = append(queues, rabbitApi.Queue{
        Name: "echo",
    })
    for _, queue := range queues {
        if utils.IsPoisonQueue(queue.Name) {
            continue
        }
        consumer, err := info.RabbitAmqp.Declare(queue.Name, info.WorkersPerQueue, events.QueuePop)
        if err != nil {
            panic(err)
        }
        go func(subConsumer rabbit.Consumer) {
            defer consumer.Close()
        loop:
            for true {
                select {
                case <-info.Shutdown:
                    break loop
                }
            }
        }(*consumer)
    }
}
