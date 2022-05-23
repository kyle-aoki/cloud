package main

import (
	"cloud/pkg/amazon"
	"cloud/pkg/syntax"
)

func main() {
	// defer util.Recover()
	amazon.InitEC2Client()
	syntax.FindAndExecute()
}
