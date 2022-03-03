package cmd

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Llongfile | log.Lmicroseconds)
}

func Execute() (err error) {
	rootCmd, err := newRootCmd()
	if err != nil {
		return
	}
	leaderBoardCmd, err := newLeaderBoardCmd()
	if err != nil {
		return
	}

	rootCmd.AddCommand(leaderBoardCmd)

	err = rootCmd.Execute()
	return
}
