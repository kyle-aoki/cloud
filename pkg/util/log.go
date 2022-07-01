package util

import (
	"cloudlab/pkg/args"
	"io"
	"log"
	"os"
)

func InitLogging() {
	if args.BoolFlag("v") {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}
