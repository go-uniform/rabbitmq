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

	cmd := _base.Command("queue.delete", func(cmd *cobra.Command, args []string) {
		service.Command("queue.delete", time.Minute, _base.NatsUri, _base.CompileNatsOptions(), map[string]string{
			"queueName": queueName,
		}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Delete an existing queue")

	cmd.Flags().StringVar(&queueName, "queueName", "", "The name/topic of the queue to be deleted")

	if err := cmd.MarkFlagRequired("queueName"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(cmd)
}
