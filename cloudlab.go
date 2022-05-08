package main

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/command"
	"cloud/pkg/config"
	"cloud/pkg/help"
	"cloud/pkg/util"
)

func main() {
	defer util.Recover()
	config.Load()
	amazon.InitEC2Client()

	switch args.Poll() {
	    case "init":           command.Initialize()
		case "destroy":        command.Destroy()
        case "delete":         command.DeleteNodes()
        case "create":         command.Create()
        case "list":           command.PrintNodes()
        case "start":          command.Start()
        case "stop":           command.Stop()
        case "config":         command.Config()
		default:               help.FatalHelpText()
	}
}
