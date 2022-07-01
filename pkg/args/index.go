package args

import (
	"fmt"
	"os"
	"strings"
)

var args []string
var flags []CharFlag

func Set() {
	rawArgs := os.Args[1:]
	cfs, argIndicies := ParseFlags(rawArgs)
	for _, index := range argIndicies {
		args = append(args, rawArgs[index])
	}
	flags = cfs
}

func BoolFlag(name string) bool {
	for _, flag := range flags {
		if flag.Name == name {
			return true
		}
	}
	return false
}

func StrFlag(name string, defaultValue string) string {
	for _, flag := range flags {
		if flag.Name == name {
			return flag.Value
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

func PollNonArgumentOrEmpty() string {
	if len(args) == 0 {
		return ""
	}
	if strings.Contains(args[0], "-") {
		fmt.Println(args)
		return ""
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
