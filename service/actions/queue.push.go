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
	var model struct {
		QueueName string `bson:"queueName"`
		Message   []byte `bson:"message"`
	}
	r.Read(&model)

	p.Notice("queue.push", diary.M{
		"model": model,
	})

	info.Rabbitmq.Push(model.QueueName)

	if r.CanReply() {
		if err := r.Reply(uniform.Request{}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
