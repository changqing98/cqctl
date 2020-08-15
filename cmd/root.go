package cmd

import (
	"github.com/changqing98/cqctl/cmd/test"
	"github.com/changqing98/cqctl/cmd/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cqctl SUBCOMMAND",
		Short: "'My go cli",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(test.NewCmdTest())
	rootCmd.AddCommand(version.NewCmdVerion())
}
