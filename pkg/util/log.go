package util

import (
	"fmt"
	"log"
)

func Log(format string, a ...any) {
	log.Println(fmt.Sprintf(format, a...))
}
