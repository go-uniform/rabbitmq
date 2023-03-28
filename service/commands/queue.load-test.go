package commands

import (
    "encoding/json"
    "fmt"
    "github.com/go-diary/diary"
    "github.com/go-uniform/uniform"
    rabbit "github.com/wagslane/go-rabbitmq"
    "service/service/_base"
    "service/service/info"
    "time"
)

const LOAD_TEST_QUEUE_NAME = "load-test"
const LOAD_TEST_MESSAGES = 10000

func init() {
    _base.Subscribe(_base.TargetCommand("queue.load-test"), loadTest)
}

func loadTest(r uniform.IRequest, p diary.IPage) {
    var uptakeCount int64 = 0
    var downtakeCount int64 = 0

    consumer, err := info.RabbitAmqp.Declare(LOAD_TEST_QUEUE_NAME, info.WorkersPerQueue, func(message rabbit.Delivery) rabbit.Action {
        downtakeCount++
        return rabbit.Ack
    })
    if err != nil {
        panic(err)
    }
    defer consumer.Close()

    startedAt := time.Now()
    for uptakeCount < LOAD_TEST_MESSAGES {
        if err := info.RabbitAmqp.Publish(LOAD_TEST_QUEUE_NAME, []byte(fmt.Sprintf("MESSAGE#%d", uptakeCount+1))); err != nil {
            panic(err)
        }
        uptakeCount++
    }
    durationSeconds := time.Now().Sub(startedAt).Seconds()
    downtake := float64(downtakeCount) / durationSeconds
    uptake := float64(uptakeCount) / durationSeconds

    output := diary.M{
        "uptakeCount":     uptakeCount,
        "downtakeCount":   downtakeCount,
        "durationSeconds": durationSeconds,
        "uptake":          uptake,
        "downtake":        downtake,
    }
    p.Notice("load-test.output", output)

    if r.CanReply() {
        data, err := json.MarshalIndent(output, "", "    ")
        if err != nil {
            panic(err)
        }
        if err := r.Reply(uniform.Request{
            Model: string(data),
        }); err != nil {
            p.Error("reply", err.Error(), diary.M{
                "err": err,
            })
        }
    }
}
