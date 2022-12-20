package args

import (
	"strings"
	"testing"
)

// program arguments
func pa(str string) []string {
	return strings.Split(str, " ")
}

func TestParseFlags(t *testing.T) {
	flags, nonFlagArgs := parseCloudlabFlags(pa("--type=t2.medium run"))
	if flags.InstanceType != "t2.medium" || len(nonFlagArgs) != 1 || nonFlagArgs[0] != "run" {
		t.Fatal(flags, nonFlagArgs)
	}
	flags, nonFlagArgs = parseCloudlabFlags(pa("run --type=t2.medium"))
	if flags.InstanceType != "t2.medium" || len(nonFlagArgs) != 1 || nonFlagArgs[0] != "run" {
		t.Fatal(flags, nonFlagArgs)
	}
	flags, nonFlagArgs = parseCloudlabFlags(pa("-q open-port -q 3000 -q i3 -q"))
	if !flags.Quiet || len(nonFlagArgs) != 3 {
		t.Fatal(flags, nonFlagArgs)
	}
	flags, nonFlagArgs = parseCloudlabFlags(pa("--all run"))
	if !flags.ShowTerminated || len(nonFlagArgs) != 1 {
		t.Fatal(flags, nonFlagArgs)
	}
	flags, nonFlagArgs = parseCloudlabFlags(pa("-a run"))
	if !flags.ShowTerminated || len(nonFlagArgs) != 1 {
		t.Fatal(flags, nonFlagArgs)
	}
}
