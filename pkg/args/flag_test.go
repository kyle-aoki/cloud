package args

import (
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	s := "subcommand -v --script=lmao.txt do something -p 3306 -d"
	input := strings.Split(s, " ")
	t.Log(input)
	flags, args := Par(input)
	t.Log(flags)
	t.Log(args)
	if flags[0].Flag != "v" {
		t.Error("flags 0 does not equal v")
	}
	if len(args) != 3 {
		t.Error("args does not equal len 3")
	}
}
