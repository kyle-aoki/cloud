package util

import (
	"fmt"
	"strings"
	"time"
)

const WaitPrintSeconds = 5

func Wait(seconds int64) {
	time.Sleep(time.Duration(seconds * int64(time.Second)))
}

func WaitPrint(message string) {
	fmt.Println(message)
	Wait(WaitPrintSeconds)
}

func WaitPrintBlock(block string) {
	lines := strings.Split(block, "\n")
	for _, line := range lines {
		WaitPrint(line)
	}
}
