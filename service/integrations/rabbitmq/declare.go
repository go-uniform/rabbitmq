package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
	"strings"
)

func (r *rabbitmq) Declare(queueName string) (amqp.Queue, <-chan amqp.Delivery, error) {
	if strings.TrimSpace(queueName) == "" {
		return amqp.Queue{}, nil, errors.New("can't create a new queue without a queueName")
	}

	queue, err := r.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return amqp.Queue{}, nil, err
	}

	queueChannel, err := r.Channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return amqp.Queue{}, nil, err
	}

	return queue, queueChannel, nil
}
