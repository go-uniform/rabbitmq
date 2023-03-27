package rabbitmq

import (
	"github.com/go-diary/diary"
	"github.com/streadway/amqp"
)

type rabbitmq struct {
	Page       diary.IPage
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (r *rabbitmq) Close() error {
	return r.Channel.Close()
}

type IRabbitmq interface {
	Push(queueName string) error
	Declare(queueName string) (amqp.Queue, <-chan amqp.Delivery, error)
	Inspect(queueName string) (amqp.Queue, error)
	Delete(queueName string) (int, error)
	Close() error
}

func NewRabbitmqConnector(page diary.IPage) IRabbitmq {
	var instance IRabbitmq
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	if err := page.Scope("rabbitmq", func(p diary.IPage) {
		instance = &rabbitmq{
			Page:       page,
			Connection: connection,
			Channel:    channel,
		}
	}); err != nil {
		panic(err)
	}

	return instance
}
