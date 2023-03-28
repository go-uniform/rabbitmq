package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"time"
)

func init() {
	var queueName string

	cmd := _base.Command("queue.create", func(cmd *cobra.Command, args []string) {
		service.Command("queue.create", time.Minute, _base.NatsUri, _base.CompileNatsOptions(), map[string]interface{}{
			"queueName": queueName,
		}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Create a new queue")

	cmd.Flags().StringVar(&queueName, "queueName", "", "The name/topic of the queue to be created")

	if err := cmd.MarkFlagRequired("queueName"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(cmd)
}
