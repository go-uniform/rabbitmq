package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/info"
)

func init() {
	_base.Subscribe(_base.TargetAction("queue", "delete"), queueDelete)
}

func queueDelete(r uniform.IRequest, p diary.IPage) {
	var model struct {
		QueueName string `bson:"queueName"`
	}
	r.Read(&model)

	p.Notice("queue.delete", diary.M{
		"model": model,
	})

	info.Rabbitmq.Declare("test")

	if r.CanReply() {
		if err := r.Reply(uniform.Request{}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
