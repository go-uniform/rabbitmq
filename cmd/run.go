package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/go-diary/diary"
	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
	"service/service"
)

var natsUri string
var natsCert string
var natsKey string
var level string
var rate int
var test bool
var disableTls bool

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run " + service.AppName + " service",
	Long:  "Run " + service.AppName + " service",
	Run: func(cmd *cobra.Command, args []string) {
		handler := diary.HumanReadableHandler
		if test {
			// test mode is used for creating an automated testing environment
			fmt.Println("entering test mode")
			handler = nil
		}
		var natsOptions = make([]nats.Option, 0)
		if !disableTls {
			cert, err := tls.LoadX509KeyPair(natsCert, natsKey)
			if err != nil {
				panic(err)
			}
			config := &tls.Config{
				// since NATS backbone for Fluid will always be on a private line we must skip host verification step
				InsecureSkipVerify: true,

				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS12,
			}
			natsOptions = append(natsOptions, nats.Secure(config))
		}
		service.Execute(test, natsUri, natsOptions, level, rate, handler, service.M{
			// todo: add custom args here
		})
	},
}

func init() {
	runCmd.Flags().StringVarP(&natsUri, "nats", "n", nats.DefaultURL, "The nats cluster URI")
	runCmd.Flags().StringVarP(&natsCert, "nats-cert", "", "/etc/ssl/certs/ssl-bundle.crt", "The nats cluster TLS certificate file path")
	runCmd.Flags().StringVarP(&natsKey, "nats-key", "", "/etc/ssl/private/ssl.key", "The nats cluster TLS key file path")
	runCmd.Flags().StringVarP(&level, "lvl", "l", "notice", "The logging level ['trace', 'debug', 'info', 'notice', 'warning', 'error', 'fatal'] that service is running in")
	runCmd.Flags().IntVarP(&rate, "rate", "r", 1000, "The sample rate of the trace logs used for performance auditing [set to -1 to log every trace]")
	runCmd.Flags().BoolVar(&test, "test", false, "A flag indicating if service should enter into test mode")
	runCmd.Flags().BoolVar(&disableTls, "disable-tls", false, "A flag indicating if service should disable tls encryption")

	rootCmd.AddCommand(runCmd)
}
