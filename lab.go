package main

import (
	"cloud/pkg/amazon"
	"cloud/pkg/syntax"
	"cloud/pkg/util"
)

func main() {
	defer util.Recover()
	amazon.InitEC2Client()
	syntax.FindAndExecute()
}
