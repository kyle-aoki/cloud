package args

import (
	"flag"
	"io"
	"log"
	"os"
)

var args []string

type IFlags struct {
	Verbose  *bool
	V        *bool
	Private  *bool
	P        *bool
	Ami      *string
	InstType *string
	Name     *string
	Script   *string
	Gigs     *string
}

var Flags IFlags

func Init() {
	Flags = IFlags{
		Verbose:  flag.Bool("verbose", false, "verbose logging"),
		V:        flag.Bool("v", false, "verbose logging"),
		Private:  flag.Bool("private", false, "create a private instance"),
		P:        flag.Bool("p", false, "create a private instance"),
		Name:     flag.String("name", "", "name of instance"),
		InstType: flag.String("type", "t2.nano", "specifiy an instance type (e.g. t2.nano)"),
		Gigs:     flag.String("gigs", "8", "number of gigabytes of storage"),
		Script:   flag.String("script", "", "path to bash script file to run on EC2 startup"),
	}
	flag.Parse()

	if *Flags.V || *Flags.Verbose {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
	args = flag.Args()
}

func IsEmpty(s string, defaulte string) string {
	if s == "" {
		return defaulte
	}
	return s
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
