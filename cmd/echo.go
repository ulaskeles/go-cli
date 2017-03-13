package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
)

var times int

var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long:  `echo is for echoing anything back.`,
	Run: echoRun,
}

func echoRun(cmd *cobra.Command, args []string) {
	for i := 0; i < times; i++ {
		fmt.Println(strings.Join(args, " "))
	}
}

func init() {
	RootCmd.AddCommand(cmdEcho)
	cmdEcho.Flags().IntVarP(&times, "times", "n", 1, "times to echo")
}