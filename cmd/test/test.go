package test

import (
	"github.com/spf13/cobra"
)

func NewCmdTest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test SUBCOMMAND",
		Short: "test short",
		Long:  "test long",
	}
	cmd.AddCommand(NewCmdTestDisplay())
	return cmd
}
