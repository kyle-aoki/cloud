package util

import (
	"log"
)

func Log(message string, a ...any) {
	log.Printf(message+"\n", a...)
}
