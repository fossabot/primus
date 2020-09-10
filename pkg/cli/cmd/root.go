package cmd

import (
	"os"

	"github.com/raba-jp/primus/pkg/cli/logging"
	"github.com/raba-jp/primus/pkg/cli/ui"
	"github.com/spf13/cobra"
)

func NewPrimusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "primus",
		Short: "provisioning tool for local machine",
	}

	cmd.AddCommand(
		NewPlanCommand(),
		NewApplyCommand(),
		NewVersionCommand(),
	)
	AddLoggingFlag(cmd)

	return cmd
}

func Execute() {
	cmd := NewPrimusCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func AddLoggingFlag(cmd *cobra.Command) {
	var logLevel string

	cmd.PersistentFlags().StringVar(&logLevel, "logLevel", "", "Set log level. Allow info, debug, warn, and error")
	cobra.OnInitialize(func() {
		if err := logging.EnableLogger(logLevel); err != nil {
			ui.Errorf("%s", err)
			os.Exit(1)
		}
	})
}
