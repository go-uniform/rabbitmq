package rabbitAmqp

import (
    "errors"
    rabbit "github.com/wagslane/go-rabbitmq"
    "service/service/utils"
    "strings"
)

func (r *rabbitAmqp) Declare(queueName string, workers int, handler rabbit.Handler) (*rabbit.Consumer, error) {
    if strings.TrimSpace(queueName) == "" {
        return nil, errors.New("can't create a new queue without a queueName")
    }

    _, err := rabbit.NewConsumer(
        r.Connection,
        nil,
        utils.GetPoisonQueue(queueName),
        rabbit.WithConsumerOptionsConcurrency(0),
        rabbit.WithConsumerOptionsQueueDurable,
        rabbit.WithConsumerOptionsConsumerName(queueName),
    )
    if err != nil {
        panic(err)
    }

    consumer, err := rabbit.NewConsumer(
        r.Connection,
        handler,
        queueName,
        rabbit.WithConsumerOptionsConcurrency(workers),
        rabbit.WithConsumerOptionsQueueDurable,
        rabbit.WithConsumerOptionsConsumerName(queueName),
    )
    if err != nil {
        panic(err)
    }

    return consumer, nil
}
