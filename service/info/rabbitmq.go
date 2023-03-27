package info

import (
	"github.com/go-diary/diary"
	"service/service/integrations/rabbitmq"
	"sync"
)

var Rabbitmq rabbitmq.IRabbitmq
var Shutdown chan bool
var Group *sync.WaitGroup
var DiaryPage diary.IPage
