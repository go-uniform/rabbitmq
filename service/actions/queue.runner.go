package actions

import (
	"github.com/go-diary/diary"
	"github.com/streadway/amqp"
	"sync"
)

func queueRunner(shutdown chan bool, group *sync.WaitGroup, p diary.IPage, queue amqp.Queue, queueChannel <-chan amqp.Delivery) {
loop:
	for true {
		select {
		case <-shutdown:
			break loop
		case message, ok := <-queueChannel:
			if ok {
				workChannel <- message
			}
		}
	}
}
