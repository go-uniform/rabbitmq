package commands

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
)

func init() {
	_base.Subscribe(_base.TargetCommand("queue.push"), push)
}

func push(r uniform.IRequest, p diary.IPage) {
	var params uniform.P
	r.Read(&params)

	if err := r.Conn().Request(p, _base.TargetAction("queue", "push"), r.Remainder(), uniform.Request{}, func(sub uniform.IRequest, _ diary.IPage) {
		if sub.HasError() {
			panic(sub.Error())
		}
		if r.CanReply() {
			var model interface{}
			sub.Read(&model)
			if err := r.Reply(uniform.Request{
				Parameters: sub.Parameters(),
				Context:    sub.Context(),
				Model:      model,
			}); err != nil {
				p.Error("reply", err.Error(), diary.M{
					"err": err,
				})
			}
		}
	}); err != nil {
		panic(err)
	}
}
