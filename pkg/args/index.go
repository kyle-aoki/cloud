package args

import (
	"cloudlab/pkg/util"
	"flag"
	"io"
	"log"
	"os"
)

type IFlags struct {
	Verbose  *bool
	V        *bool
	Private  *bool
	P        *bool
	InstType *string
	Name     *string
	Script   *string
	Gigs     *string
	All      *bool
	Quiet    *bool
}

var Flags IFlags
var Args []string

func Init() {
	Flags = IFlags{
		Verbose:  flag.Bool("verbose", false, "verbose logging"),
		V:        flag.Bool("v", false, "verbose logging"),
		Private:  flag.Bool("private", false, "(lab run, lab ssh) select private instance"),
		P:        flag.Bool("p", false, "(lab run, lab ssh) select private instance"),
		Name:     flag.String("name", "", "(lab run) name of instance"),
		InstType: flag.String("type", "t2.nano", "(lab run) specifiy an instance type"),
		Gigs:     flag.String("gigs", "8", "(lab run) number of gigabytes of storage"),
		Script:   flag.String("script", "", "(lab run) path to bash script file to run on EC2 startup"),
		All:      flag.Bool("all", false, "(lab list, lab watch) show terminated instances"),
		Quiet:    flag.Bool("q", false, "(lab list) list names only"),
	}
	flag.Parse()

	if *Flags.V || *Flags.Verbose {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
	Args = flag.Args()
	util.Log("program args: %v", Args)
}

func FlagExists(s *string) bool {
	return s != nil && *s != ""
}

func Poll() string {
	if len(Args) == 0 {
		panic("not enough arguments")
	}
	next := Args[0]
	Args = Args[1:]
	return next
}

func PollOrEmpty() string {
	if len(Args) == 0 {
		return ""
	}
	next := Args[0]
	Args = Args[1:]
	return next
}

func Collect() []string {
	if len(Args) == 0 {
		panic("not enough arguments")
	}
	collection := Args[:]
	Args = nil
	return collection
}

func CollectOrEmpty() []string {
	collection := Args[:]
	Args = nil
	return collection
}
