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
	//add functions
}

func NewRabbitmqConnector(page diary.IPage) IRabbitmq {
	var instance IRabbitmq
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
