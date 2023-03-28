package rabbitAmqp

import (
    "github.com/go-diary/diary"
    rabbit "github.com/wagslane/go-rabbitmq"
)

type rabbitAmqp struct {
    Page       diary.IPage
    Connection *rabbit.Conn
    Publisher  *rabbit.Publisher
}

func (r *rabbitAmqp) Close() error {
    return r.Connection.Close()
}

type IRabbitAmqp interface {
    Publish(queueName string, message []byte) error
    Declare(queueName string, workers int, handler rabbit.Handler) (*rabbit.Consumer, error)
    Close() error
}

func NewRabbitAmqpConnector(page diary.IPage, uri string, config rabbit.Config) IRabbitAmqp {
    connection, err := rabbit.NewConn(uri, rabbit.WithConnectionOptionsConfig(config))
    if err != nil {
        panic(err)
    }

    publisher, err := rabbit.NewPublisher(connection)
    if err != nil {
        panic(err)
    }

    var instance IRabbitAmqp
    if err := page.Scope("rabbitAmqp", func(p diary.IPage) {
        instance = &rabbitAmqp{
            Page:       page,
            Connection: connection,
            Publisher:  publisher,
        }
    }); err != nil {
        panic(err)
    }

    return instance
}
