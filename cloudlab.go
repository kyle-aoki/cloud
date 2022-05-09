package main

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/command"
	"cloud/pkg/config"
	"cloud/pkg/help"
	"cloud/pkg/util"
)

var CommandMap = map[string]func(){
	"init":    command.Initialize,
	"destroy": command.Destroy,
	"delete":  command.DeleteNodes,
	"list": func() {
		exec(map[string]func(){
			"":     command.PrintNodes,
			"keys": command.ListKeys,
		})
	},
	"create": func() {
		exec(map[string]func(){
			"instance": command.CreateInstance,
			"key":      command.CreateKeyPair,
		})
	},
	"start":  command.Start,
	"stop":   command.Stop,
	"config": command.Config,
}

func main() {
	defer util.Recover()
	config.Load()
	amazon.InitEC2Client()

	exec(CommandMap)
}

func exec(commandMap map[string]func()) {
	if command, ok := commandMap[args.PollOrEmpty()]; ok {
		command()
	} else {
		help.FatalHelpText()
	}
}
