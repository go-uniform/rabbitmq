package cmd

import (
	"github.com/spf13/cobra"
	"service/service"
)

var exampleOneCmd = &cobra.Command{
	Use:   "command:example-one",
	Short: "Request the running " + service.AppName + " to execute the example-one command",
	Long:  "Request the running " + service.AppName + " to execute the example-one command",
	Run: func(cmd *cobra.Command, args []string) {
		service.InitializeDiary(test, level, rate)
		service.Command("example-one", natsUri, compileNatsOptions(), map[string]string{})
	},
}

func init() {
	rootCmd.AddCommand(exampleOneCmd)
}