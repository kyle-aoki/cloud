package args

import (
	"flag"
)

var Verbose bool

func FlagParse() {
	Verbose = *flag.Bool("v", true, "verbose")
	flag.Parse()
}
