package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
	"strings"
)

func (r *rabbitmq) Inspect(queueName string) (amqp.Queue, error) {
	if strings.TrimSpace(queueName) == "" {
		return amqp.Queue{}, errors.New("can't inspect a queue without a queueName")
	}

	queue, err := r.Channel.QueueInspect(queueName)
	if err != nil {
		return amqp.Queue{}, err
	}

	return queue, nil
}
