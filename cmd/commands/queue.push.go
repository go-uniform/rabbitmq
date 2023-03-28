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

    cmd.Flags().StringVar(&queueName, "queueName", "echo", "The name/topic of the queue to be pushed to")
    cmd.Flags().StringVar(&message, "message", "Toto, I've got a feeling we're not in Kansas anymore.", "The message to be pushed onto the given queue")

    _base.RootCmd.AddCommand(cmd)
}
