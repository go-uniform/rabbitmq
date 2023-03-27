package actions

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/streadway/amqp"
	"sync"
)

var workChannel = make(chan amqp.Delivery)

func queueWorker(shutdown chan bool, group *sync.WaitGroup, p diary.IPage, queue amqp.Queue, queueChannel <-chan amqp.Delivery) {
loop:
	for true {
		select {
		case <-shutdown:
			break loop
		case message, ok := <-workChannel:
			if ok {
				message.Ack(true)
				fmt.Sprintf("#### Received a message: %s\n", message.Body)
			}
		}
		fmt.Sprintf("#### no messages on queue [%s]\n", queue)
	}
}
