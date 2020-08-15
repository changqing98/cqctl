package grpc

import (
	"github.com/spf13/cobra"
	"os/exec"
)

var (
	lan string
	src string
	dst string
)

func NewCmdGRPC() *cobra.Command{
	cmd := &cobra.Command{
		Use: "grpc SUBCOMMAND",
		Short: "gRPC command",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}

	cmd.Flags().StringVarP(&lan, "language", "l", "all", "Output language")
	cmd.Flags().StringVarP(&src, "src", "s", "", "Source dir")
	cmd.Flags().StringVarP(&dst, "dst", "d", "./", "Destination dir")

	return cmd
}

func run()  {
	cmdStr := "protoc -I "
	exec.Command(cmdStr)
}
