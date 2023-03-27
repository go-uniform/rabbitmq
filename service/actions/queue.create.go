package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
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

	queue, queueChannel, err := info.Rabbitmq.Declare(model.QueueName)
	if err != nil {
		panic(err)
	}
	go queueRunner(info.Shutdown, info.Group, info.DiaryPage, queue, queueChannel)

	if r.CanReply() {
		if err := r.Reply(uniform.Request{}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
