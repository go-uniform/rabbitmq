package events

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	rabbit "github.com/wagslane/go-rabbitmq"
	"service/service/_base"
	"service/service/info"
	"service/service/utils"
	"strings"
	"time"
)

func QueuePop(message rabbit.Delivery) (action rabbit.Action) {
	category := _base.TargetEvent("queue", message.RoutingKey)

	if err := info.DiaryPage.Scope(category, func(p diary.IPage) {
		request, err := uniform.ParseRequest(message.Body)
		if err != nil {
			panic(err)
		}
		if request.Context == nil {
			request.Context = make(uniform.M)
		}
		message.Body = nil
		request.Context["messageDetails"] = message

		if err := info.Conn.Request(
			info.DiaryPage,
			category,
			time.Minute,
			request,
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
