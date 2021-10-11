package rabbitmq

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/streadway/amqp"
)

type rabbitmq struct {
	Page    diary.IPage
	Channel *amqp.Channel
}

type IRabbitmq interface {
	PushQueue()
	PopQueue()
}

func NewRabbitmqConnector(page diary.IPage) IRabbitmq {
	var instance IRabbitmq
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	page.Scope("rabbitmq", func(p diary.IPage) {
		instance = &rabbitmq{
			Page:    page,
			Channel: channel,
		}
	})
	return instance
}
