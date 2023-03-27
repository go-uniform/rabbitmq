package rabbitmq

import "github.com/streadway/amqp"

func (r *rabbitmq) Push(queueName string) error {
	return r.Channel.Publish(
		"",
		queueName,
		false,
		true,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		})
}
