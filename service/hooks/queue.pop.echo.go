package hooks

import (
    "fmt"
    "github.com/go-diary/diary"
    "github.com/go-uniform/uniform"
    "service/service/_base"
    "strings"
)

func init() {
    _base.Subscribe(_base.TargetEvent("queue", "pop.echo"), queuePopEcho)
}

func queuePopEcho(r uniform.IRequest, p diary.IPage) {
    var model struct {
        MessageId string `bson:"messageId"`
        Body      []byte `bson:"body"`
    }
    r.Read(&model)

    message := string(model.Body)
    if strings.HasPrefix(strings.ToLower(message), "poison") {
        panic("poison pill!")
    }

    fmt.Println()
    fmt.Println()
    fmt.Println()
    fmt.Printf("echo: %s\n", string(model.Body))
    fmt.Println()
    fmt.Println()
    fmt.Println()

    if r.CanReply() {
        if err := r.Reply(uniform.Request{}); err != nil {
            p.Error("reply", err.Error(), diary.M{
                "err": err,
            })
        }
    }
}
