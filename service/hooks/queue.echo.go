package hooks

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"strings"
)

func init() {
	_base.Subscribe(_base.TargetEvent("queue", "echo"), queueEcho)
}

func queueEcho(r uniform.IRequest, p diary.IPage) {
	var model []byte
	r.Read(&model)

	message := string(model)
	if strings.HasPrefix(strings.ToLower(message), "poison") {
		panic("poison pill!")
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("echo: %s\n", message)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	if r.CanReply() {
		if err := r.Reply(uniform.Request{}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
