package cmd

import (
	"github.com/go-uniform/uniform"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"service/service/info"
)

func init() {
	var level string
	var rate int
	var limit int
	var test bool
	var virtual bool
	var workers int

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run " + info.AppName + " service",
		Long:  "Run " + info.AppName + " service",
		Run: func(cmd *cobra.Command, args []string) {
			service.Execute(level, rate, limit, test, virtual, _base.NatsUri, _base.CompileNatsOptions(), uniform.M{
				"nats":       _base.NatsUri,
				"natsCert":   _base.NatsCert,
				"natsKey":    _base.NatsKey,
				"disableTls": _base.DisableTls,
				"lvl":        level,
				"rate":       rate,
				"limit":      limit,
				"test":       test,
				"virtual":    virtual,
				"workers":    workers,
			})
		},
	}

	// set the service's environment configurations via many command-line-interface (CLI) arguments
	runCmd.Flags().StringVarP(&level, "lvl", "l", "notice", "The logging level ['trace', 'debug', 'info', 'notice', 'warning', 'error', 'fatal'] that service is running in")
	runCmd.Flags().IntVarP(&rate, "rate", "r", 1000, "The sample rate of the trace logs used for performance auditing [set to -1 to log every trace]")
	runCmd.Flags().IntVarP(&limit, "limit", "x", 1000, "The messages per second that each topic worker will be limited to [set to 0 or less for maximum throughput]")
	runCmd.Flags().BoolVar(&test, "test", false, "A flag indicating if service should enter into test mode")
	runCmd.Flags().BoolVar(&virtual, "virtual", false, "A flag indicating if service should virtualize external integration calls")
	runCmd.Flags().IntVar(&workers, "workers", 100, "The number of worker threads that should be made available for message processing")

	_base.RootCmd.AddCommand(runCmd)
}
