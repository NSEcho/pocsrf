package main

import (
	"github.com/lateralusd/pocsrf/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
