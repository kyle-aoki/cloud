package syntax

import (
	"cloudlab/pkg/args"
)

func FindAndExecute() {
	traverse(SyntaxTree)
}

func traverse(commandMap map[string]any) {
	if val, ok := commandMap[args.PollOrEmpty()]; ok {
		switch val.(type) {
		case Cmd:
			val.(Cmd).fn()
		default:
			traverse(val.(map[string]any))
		}
	} else {
		HelpText()
	}
}
