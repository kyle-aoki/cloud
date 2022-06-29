package main

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/syntax"
	"cloudlab/pkg/util"
)

func main() {
	defer util.Recover()
	args.Prepare()
	util.InitLogging()
	amazon.InitEC2Client()
	syntax.FindAndExecute()
}
