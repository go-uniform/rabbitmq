package actions

import (
    "github.com/go-diary/diary"
    "github.com/go-uniform/uniform"
    "go.mongodb.org/mongo-driver/bson"
    "service/service/_base"
    "service/service/info"
)

func init() {
    _base.Subscribe(_base.TargetAction("queue", "push"), queuePush)
}

func queuePush(r uniform.IRequest, p diary.IPage) {
    params := r.Parameters()
    queueName := params["queueName"]
    if queueName == "" {
        panic("queueName may not be empty")
    }

    var model interface{}
    r.Read(&model)

    p.Notice("queue.push", diary.M{
        "params": params,
        "model":  model,
    })

    payload, err := bson.Marshal(model)
    if err != nil {
        panic(err)
    }

    if err := info.RabbitAmqp.Publish(params["queueName"], payload); err != nil {
        panic(err)
    }

    if r.CanReply() {
        if err := r.Reply(uniform.Request{}); err != nil {
            p.Error("reply", err.Error(), diary.M{
                "err": err,
            })
        }
    }
}
