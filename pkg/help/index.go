package help

import (
	"fmt"
	"os"
)

const HelpText = `
List  Nodes:        cloud list
Create Node:        cloud create <launch-template> <name-tag-1> <name-tag-2> ...
Delete Node:        cloud delete <node-name-tag or instance-id-substring> ...

`

func FatalHelpText() {
	fmt.Print(HelpText)
	os.Exit(1)
}
