package args

import (
	"flag"
	"fmt"
	"strings"
)

var args []string
var Verbose bool

func Prepare() {
	v := flag.Bool("v", false, "verbose")
	
	flag.Parse()
	fmt.Println("v", *v)
	args = flag.Args()
	fmt.Println(args)
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
