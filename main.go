package main

import (
	"main/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
