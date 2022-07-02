package args

import (
	"io"
	"log"
	"os"
)

var args []string
var flags []Flag

func Prepare() {
	rawArgs := os.Args[1:]
	flgs, argIndicies := ExtractFlagsFromArgs(rawArgs)
	for _, index := range argIndicies {
		args = append(args, rawArgs[index])
	}
	flags = flgs
}

func InitLogging() {
	if BoolFlag("v", "verbose") {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}

func BoolFlag(names ...string) bool {
	for _, flag := range flags {
		for _, name := range names {
			if flag.Name == name {
				return true
			}
		}
	}
	return false
}

func StrFlag(defaultValue string, names ...string) string {
	for _, flag := range flags {
		for _, name := range names {
			if flag.Name == name {
				return flag.Value
			}
		}
	}
	return defaultValue
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
