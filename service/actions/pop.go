package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/info"
)

func init() {
	_base.Subscribe(_base.TargetAction("queue", "pop"), popQueue)
}

func popQueue(r uniform.IRequest, p diary.IPage) {

	p.Notice("queue.pop", nil)

	info.Rabbitmq.PopQueue()

	if r.CanReply() {
		if err := r.Reply(uniform.Request{}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
