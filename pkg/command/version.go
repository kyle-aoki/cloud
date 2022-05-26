package command

import "fmt"

const version = "1.0.0"

func Version() {
	fmt.Printf("cloudlab v%v\n", version)
}
