package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"time"
)

func init() {

	cmd := _base.Command("queue.load-test", func(cmd *cobra.Command, args []string) {
		service.Command("queue.load-test", time.Minute*10, _base.NatsUri, _base.CompileNatsOptions(), map[string]string{}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Load test the queue service")

	_base.RootCmd.AddCommand(cmd)
}
