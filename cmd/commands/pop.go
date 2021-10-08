package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"time"
)

func init() {

	cmd := _base.Command("pop", func(cmd *cobra.Command, args []string) {
		service.Command("pop", time.Minute, _base.NatsUri, _base.CompileNatsOptions(), map[string]string{}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Push a message to a queue")

	_base.RootCmd.AddCommand(cmd)
}
