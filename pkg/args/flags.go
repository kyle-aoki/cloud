package args

import (
	"os"
)

type CloudlabFlags struct {
	Verbose        bool
	Quiet          bool
	Private        bool
	ShowTerminated bool
	InstanceType   string
	Gigabytes      string
	InstanceName   *string
}

func cloudlabFlagDefaults() *CloudlabFlags {
	return &CloudlabFlags{
		Verbose:        false,
		Quiet:          false,
		Private:        false,
		ShowTerminated: false,
		InstanceType:   "t2.micro",
		Gigabytes:      "8gb",
		InstanceName:   nil,
	}
}

var Flags *CloudlabFlags
var programArgs []string

func ParseProgramInput() {
	rawArgs := os.Args[1:]
	flags, args := parseCloudlabFlags(rawArgs)
	programArgs = args
	Flags = flags
}
