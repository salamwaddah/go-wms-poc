package main

import (
	"go-wms-poc/cmd"
	"go-wms-poc/config"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		panic("error creating config")
	}

	cmd.Execute()
}
