package command

import "fmt"

const version = "1.0.0"

func Version() {
	fmt.Printf("cloudlab version v%v\n", version)
}
