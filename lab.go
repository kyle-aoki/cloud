package main

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/cmd"
	"cloudlab/pkg/util"
	"fmt"
)

func main() {
	defer util.MainRecover()
	args.Prepare()
	args.InitLogging()
	amazon.InitEC2Client()

	if command, ok := Syntax[args.PollOrEmpty()]; ok {
		command()
	} else {
		fmt.Println("help text")
	}
}

var Syntax = map[string]func(){
	"version":    cmd.Version,
	"info":       cmd.Info,
	"init":       cmd.InitializeCloudLabResources,
	"destroy":    cmd.DestroyCloudLabResources,
	"list":       cmd.ListInstances,
	"run":        cmd.CreateInstance,
	"delete":     cmd.DeleteInstances,
	"open-port":  cmd.OpenPort,
	"close-port": cmd.ClosePort,
}
