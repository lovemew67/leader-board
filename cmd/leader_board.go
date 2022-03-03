package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func newLeaderBoardCmd() (result *cobra.Command, err error) {
	result = &cobra.Command{
		Use: "leader-board",
		Run: func(cmd *cobra.Command, args []string) {
			funcName := "leaderBoardCmd.Run"
			log.Printf("[%s] in run function", funcName)
		},
	}
	return
}
