package main

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/cmd"
	"cloudlab/pkg/util"
	"fmt"
)

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

func main() {
	defer util.MainRecover()
	args.Init()
	amazon.InitEC2Client()

	if command, found := Syntax[args.PollOrEmpty()]; found {
		command()
	} else {
		fmt.Println("help text")
	}
}
