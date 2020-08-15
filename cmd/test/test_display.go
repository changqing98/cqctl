package test

import (
	"fmt"

	"github.com/spf13/cobra"
)

var(
	name string
)

func initOptions(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&name, "message", "m", "hello", "Default value hello")
}

func NewCmdTestDisplay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "display MESSAGE",
		Short: "Test short",
		Long:  "Test display long",
		Run: func(cmd *cobra.Command, args []string) {
			runTestDisplay()
		},
	}
	initOptions(cmd)
	return cmd
}

func runTestDisplay() {
	fmt.Printf("Test display %s  ...\n", name)
}
