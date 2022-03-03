package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func newRootCmd() (result *cobra.Command, err error) {
	result = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			funcName := "rootCmd.Run"
			log.Printf("[%s] in run function", funcName)
		},
	}
	return
}
