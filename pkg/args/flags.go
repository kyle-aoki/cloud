package args

import (
	"os"
)

type CloudlabFlags struct {
	PrintVersion   bool
	Verbose        bool
	Quiet          bool
	ShowHelp       bool
	Private        bool
	ShowTerminated bool
	InstanceType   string
	Gigabytes      string
	InstanceName   *string
}

func cloudlabFlagDefaults() *CloudlabFlags {
	return &CloudlabFlags{
		InstanceType: "t2.micro",
		Gigabytes:    "8",
		InstanceName: nil,
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
