package util

import (
	"fmt"
	"os"
)

func Exit(message string) {
	fmt.Println(message)
	os.Exit(0)
}
