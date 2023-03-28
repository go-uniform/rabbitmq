package service

import (
    "github.com/go-diary/diary"
    "service/service/info"
    "sync"
)

func ShutdownAfter(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
    info.RabbitAmqp.Close()
}
