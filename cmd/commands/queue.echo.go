package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"time"
)

func init() {
	var message string

	cmd := _base.Command("queue.echo", func(cmd *cobra.Command, args []string) {
		service.Command("queue.echo", time.Minute, _base.NatsUri, _base.CompileNatsOptions(), map[string]interface{}{
			"message": message,
		}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Push message to the echo queue")

	cmd.Flags().StringVar(&message, "message", "Toto, I've got a feeling we're not in Kansas anymore.", "The message to be pushed onto the given queue")

	_base.RootCmd.AddCommand(cmd)
}
