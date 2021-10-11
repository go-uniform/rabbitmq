package service

import (
	"github.com/go-diary/diary"
	"service/service/info"
	"service/service/integrations/rabbitmq"
	"sync"
)

func RunBefore(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
	info.Rabbitmq = rabbitmq.NewRabbitmqConnector(p)
}
