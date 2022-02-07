package main

import (
	"cloud/pkg/args"
	"cloud/pkg/command"
	"cloud/pkg/help"
	"cloud/pkg/util"
)

func main() {
	defer util.Recover()

	switch args.Poll() {
		case "delete":         command.DeleteNodes()
		case "create":         command.Create()
		case "list":           command.PrintNodes()
        case "start":          command.Start()
        case "stop":           command.Stop()
		default:               help.Print()
	}
}
