package events

import (
    "github.com/go-diary/diary"
    "github.com/go-uniform/uniform"
    rabbit "github.com/wagslane/go-rabbitmq"
    "go.mongodb.org/mongo-driver/bson"
    "service/service/_base"
    "service/service/info"
    "service/service/utils"
    "strings"
    "time"
)

func QueuePop(message rabbit.Delivery) (action rabbit.Action) {
    var model interface{}
    if err := bson.Unmarshal(message.Body, &model); err != nil {
        panic(err)
    }

    category := _base.TargetEvent("queue", message.RoutingKey)
    if err := info.DiaryPage.Scope(category, func(p diary.IPage) {
        if err := info.Conn.Request(
            info.DiaryPage,
            category,
            time.Minute,
            uniform.Request{
                Parameters: map[string]string{
                    "messageId": message.MessageId,
                },
                Model: model,
            },
            func(sub uniform.IRequest, _ diary.IPage) {
                if sub.HasError() {
                    panic(sub.Error())
                }
            },
        ); err != nil {
            panic(err)
        }
    }); err != nil {
        if strings.HasPrefix(err.Error(), "nats: timeout") {
            return rabbit.NackRequeue
        }
        if err := info.RabbitAmqp.Publish(
            utils.GetPoisonQueue(message.RoutingKey),
            message.Body,
        ); err != nil {
            panic(err)
        }
        return rabbit.NackDiscard
    }
    return rabbit.Ack
}
