package help

import (
	"fmt"
	"os"
)

const HelpText = `
    Show Config:        cloud config
    List  Nodes:        cloud list
    Create Node:        cloud create <config-type> <name-tag-1> <name-tag-2> ...
    Delete Node:        cloud delete <node-name-tag or instance-id-substring> ...

`

func Print() {
	fmt.Print(HelpText)
	os.Exit(1)
}
