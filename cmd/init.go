package cmd

import (
	"encoding/json"
	"log"
	"strings"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileDefaultName = ".leaderboard"
)

var (
	configFilePath = ""
)

func initConfig() {
	funcName := "initConfig"

	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
		log.Printf("[%s] set config file path: %s", funcName, configFilePath)
	} else {
		hd, err := homedir.Dir()
		if err != nil {
			log.Panicf("[%s] failed to get homedir, err: %+v", funcName, err)
		}
		viper.AddConfigPath(hd)
		viper.SetConfigName(configFileDefaultName)
		log.Printf("[%s] set to %s/%s.(json/toml/yaml/yml/properties/props/prop/hcl/tfvars/dotenv/env/ini)", funcName, hd, configFileDefaultName)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("[%s] failed to read in config, err: %+v", funcName, err)
	}
	log.Printf("[%s] read in config", funcName)

	jsIndent, err := json.MarshalIndent(viper.AllSettings(), "", "\t")
	if err != nil {
		log.Panicf("[%s] failed to marshal with indent, err: %+v", funcName, err)
	}
	log.Printf("[%s] log config before initializing: %s", funcName, jsIndent)
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Llongfile | log.Lmicroseconds)
	cobra.OnInitialize(initConfig)
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
