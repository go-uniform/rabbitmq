package actions

import (
    "github.com/go-diary/diary"
    "github.com/go-uniform/uniform"
    rabbit "github.com/wagslane/go-rabbitmq"
    "service/service/_base"
    "service/service/events"
    "service/service/info"
)

func init() {
    _base.Subscribe(_base.TargetAction("queue", "create"), queueCreate)
}

func queueCreate(r uniform.IRequest, p diary.IPage) {
    var model struct {
        QueueName string `bson:"queueName"`
    }
    r.Read(&model)

    p.Notice("queue.create", diary.M{
        "model": model,
    })

    consumer, err := info.RabbitAmqp.Declare(model.QueueName, info.WorkersPerQueue, events.QueuePop)
    if err != nil {
        panic(err)
    }
    go func(subConsumer rabbit.Consumer) {
        defer consumer.Close()
    loop:
        for true {
            select {
            case <-info.Shutdown:
                break loop
            }
        }
    }(*consumer)

    if r.CanReply() {
        if err := r.Reply(uniform.Request{}); err != nil {
            p.Error("reply", err.Error(), diary.M{
                "err": err,
            })
        }
    }
}
