package main

import (
	"cloud/pkg/args"
	"cloud/pkg/command"
	"cloud/pkg/config"
	"cloud/pkg/help"
	"cloud/pkg/util"
)

func main() {
	defer util.Recover()
	
	config.Load()

	switch args.Poll() {
		case "config":         command.ShowConfig()
		case "delete":         command.DeleteNodes()
		case "create":         command.CreateNode()
		case "list":           command.PrintNodes()
		default:               help.Print()
	}
}
