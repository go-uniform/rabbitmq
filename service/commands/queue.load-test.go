package commands

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/streadway/amqp"
	"service/service/_base"
	"service/service/info"
	"time"
)

func init() {
	_base.Subscribe(_base.TargetCommand("queue.load-test"), loadTest)
}

func loadTestUptake(r uniform.IRequest, p diary.IPage, durationSeconds int) float64 {
	var counter = 0
	var timeLimit = time.Tick(time.Second * time.Duration(durationSeconds))

loop:
	for true {
		select {
		case <-timeLimit:
			break loop
		default:
			if err := r.Conn().Request(p, _base.TargetAction("queue", "push"), r.Remainder(), uniform.Request{
				Model: map[string]interface{}{
					"queueName": "load-test",
					"message":   fmt.Sprintf("LOADTEST#%5d", counter),
				},
			}, func(sub uniform.IRequest, _ diary.IPage) {
				if sub.HasError() {
					panic(sub.Error())
				}
			}); err != nil {
				panic(err)
			}
			break
		}
		counter++
	}

	return float64(counter) / float64(durationSeconds)
}

func loadTestDowntake(queueChannel <-chan amqp.Delivery) float64 {
	var counter = 0
	var timeLimit = time.Tick(time.Second)
	var limiter = time.NewTicker(time.Millisecond)

loop:
	for true {
		select {
		case <-timeLimit:
			break loop
		case message, ok := <-queueChannel:
			if !ok {
				<-limiter.C
				continue
			}
			message.Ack(false)
		}
		counter++
	}

	return float64(counter)
}

func loadTest(r uniform.IRequest, p diary.IPage) {
	if _, err := info.Rabbitmq.Delete("load-test"); err != nil {
		panic(err)
	}
	_, queueChannel, err := info.Rabbitmq.Declare("load-test")
	if err != nil {
		panic(err)
	}

	uptake := loadTestUptake(r, p, 60)
	downtake := loadTestDowntake(queueChannel)
	p.Notice("load-test.output", diary.M{
		"uptake":   uptake,
		"downtake": downtake,
	})

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: fmt.Sprintf("Uptake: %0.2f msgs/s\nDowntake: %0.2f msgs/s\n", uptake, downtake),
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
