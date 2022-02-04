package args

import (
	"cloud/pkg/help"
	"os"
)

var ExecutionPath string
var args []string

func init() {
	args = os.Args
	ExecutionPath = Poll()
}

func Poll() string {
	if len(args) == 0 {
		help.Print()
	}
	next := args[0]
	args = args[1:]
	return next
}

func Collect() []string {
	if len(args) == 0 {
		panic("Not enough arguments. Try 'cloud --help'")
	}
	return args
}
