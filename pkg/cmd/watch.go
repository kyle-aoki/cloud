package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func Watch() {
	WatchLoop(0)
	WatchLoop(0)
	WatchLoop(1)
	WatchLoop(1)
	for i := 0; i < 500; i++ {
		WatchLoop(2)
	}
	for {
		WatchLoop(3)
	}
}

func WatchLoop(interval int) {
	co := resource.New()

	var instances []*ec2.Instance
	if *args.Flags.All {
		instances = co.Finder.FindInstances()
	} else {
		instances = co.Finder.FindNonTerminatedInstances()
	}

	time.Sleep(time.Second * time.Duration(interval))
	util.ClearTerminal()
	PrintInstanceList(instances)
}
