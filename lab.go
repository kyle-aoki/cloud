package main

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/syntax"
	"cloudlab/pkg/util"
)

func main() {
	defer util.Recover()
	amazon.InitEC2Client()
	syntax.FindAndExecute()
}
