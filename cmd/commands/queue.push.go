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
	var message string

	cmd := _base.Command("queue.push", func(cmd *cobra.Command, args []string) {
		service.Command("queue.push", time.Minute, _base.NatsUri, _base.CompileNatsOptions(), map[string]string{
			"queueName": queueName,
			"message":   message,
		}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Push a message to a queue")

	cmd.Flags().StringVar(&queueName, "queueName", "", "The name/topic of the queue to be pushed to")
	cmd.Flags().StringVar(&message, "message", "", "The message to be pushed onto the given queue")

	if err := cmd.MarkFlagRequired("queueName"); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired("message"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(cmd)
}
