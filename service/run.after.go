package service

import (
	"github.com/go-diary/diary"
	"sync"
)

func RunAfter(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
	//var queueNames = []string{ "test" }
	//for _, queueName := range queueNames {
	//	queue, queueChannel, err := info.Rabbitmq.Declare(queueName)
	//	if err != nil {
	//		panic(err)
	//	}
	//	go queueRunner(shutdown, group, p, queue, queueChannel)
	//}
}
