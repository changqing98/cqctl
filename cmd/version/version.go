package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	VERSION = "1.0.0"
)

func NewCmdVerion() *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(VERSION)
		},
	}
	return cmd
}
