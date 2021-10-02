package service

import (
	"fmt"
	"github.com/go-diary/diary"
	"service/service/info"
)

func RunBefore(p diary.IPage) {
	// todo: add your custom startup logic here
	fmt.Println(info.Args)
}