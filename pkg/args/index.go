package args

import (
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
		panic("not enough arguments")
	}
	next := args[0]
	args = args[1:]
	return next
}

func PollOrEmpty() string {
	if len(args) == 0 {
		return ""
	}
	next := args[0]
	args = args[1:]
	return next
}

func Collect() []string {
	if len(args) == 0 {
		panic("not enough arguments")
	}
	collection := args[:]
	args = nil
	return collection
}
