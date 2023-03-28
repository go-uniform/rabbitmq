package rabbitAmqp

import (
    rabbit "github.com/wagslane/go-rabbitmq"
)

func (r *rabbitAmqp) Publish(queueName string, message []byte) error {
    err := r.Publisher.Publish(
        message,
        []string{queueName},
        rabbit.WithPublishOptionsPersistentDelivery,
        rabbit.WithPublishOptionsMandatory,
    )
    return err
}
