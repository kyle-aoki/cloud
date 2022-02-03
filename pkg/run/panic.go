package run

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
