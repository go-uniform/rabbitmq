package service

import (
	"github.com/go-diary/diary"
	"service/service/info"
	"service/service/integrations/rabbitmq"
	"sync"
)

func RunBefore(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
	if info.Virtualize {
		panic("this service can't be started in virtual mode, it is not yet supported")
	}
	info.Rabbitmq = rabbitmq.NewRabbitmqConnector(p)
	info.Shutdown = shutdown
	info.Group = group
	info.DiaryPage = p
}
