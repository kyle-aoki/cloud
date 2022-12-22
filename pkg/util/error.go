package util

import "fmt"

func MainRecover() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
