package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var queueChannels = make(map[string]<-chan amqp.Delivery)
var lock = sync.Mutex{}

func consume(channel *amqp.Channel, queueName string) <-chan amqp.Delivery {
	lock.Lock()
	defer lock.Unlock()

	if queueChannel, exists := queueChannels[queueName]; exists {
		return queueChannel
	}
	queueChannel, err := channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	queueChannels[queueName] = queueChannel

	return queueChannel
}

func (r *rabbitmq) Pop() string {

	queue := "hello"

	select {
	case <-time.Tick(time.Second):
		break
	case message, ok := <-consume(r.Channel, queue):
		if ok {
			defer func(m amqp.Delivery) {
				_ = m.Ack(true)
			}(message)
			return fmt.Sprintf("#### Received a message: %s\n", message.Body)
		}
	}

	return fmt.Sprintf("#### no messages on queue [%s]\n", queue)
}
