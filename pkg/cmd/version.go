package cmd

import "fmt"

const version = "1.0.0"

func Version() {
	fmt.Printf("%v\n", version)
}
