package main

import (
	"cloud/pkg/args"
	"cloud/pkg/command"
	"cloud/pkg/help"
	"cloud/pkg/util"
	"cloud/pkg/config"
)

func main() {
	defer util.Recover()
	config.Load()

	switch args.Poll() {
        case "delete":         command.DeleteNodes()
        case "create":         command.Create()
        case "list":           command.PrintNodes()
        case "start":          command.Start()
        case "stop":           command.Stop()
        case "config":         command.ShowConfig()
		default:               help.Print()
	}
}
