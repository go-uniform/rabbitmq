package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
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

	p.Notice("queue.push", diary.M{
		"params": params,
		"model":  r.Raw(),
	})

	if err := info.RabbitAmqp.Publish(params["queueName"], r.Bytes()); err != nil {
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
