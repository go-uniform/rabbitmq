package info

import (
    "fmt"
    "github.com/go-diary/diary"
    "service/service/integrations/rabbitAmqp"
    "service/service/integrations/rabbitApi"
    "sync"
)

var RabbitApi rabbitApi.IRabbitApi
var RabbitAmqp rabbitAmqp.IRabbitAmqp
var Shutdown chan bool
var Group *sync.WaitGroup
var DiaryPage diary.IPage
var WorkersPerQueue int

var GetPoisonQueue = func(queueName string) string {
    return fmt.Sprintf("_poison.%s", queueName)
}
