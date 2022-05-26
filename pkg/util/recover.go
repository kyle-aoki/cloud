package util

import "fmt"

const Ignore bool = true

func Recover() {
	if Ignore {
		return
	}
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
