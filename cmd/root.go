package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

const (
	configFilePathFlag      = "config"
	configFilePathShorthand = "c"
	configFilePathDefault   = ""
	configFilePathUsage     = "specify config file path"
)

func newRootCmd() (result *cobra.Command, err error) {
	result = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			funcName := "rootCmd.Run"
			log.Printf("[%s] in run function", funcName)
		},
	}
	result.PersistentFlags().StringVarP(&configFilePath, configFilePathFlag, configFilePathShorthand, configFilePathDefault, configFilePathUsage)
	return
}
