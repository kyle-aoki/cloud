package util

import "fmt"

// ########################################
const messageLength = 70

func VPrint(message string, value string) {
	if len(message) > messageLength {
		fmt.Println(fmt.Sprintf("%s %s", message, value))
		return
	}
	spaces := messageLength - len(message)
	fmt.Println(fmt.Sprintf("%s%s%s", message, CreateSpacerString(spaces), value))
}

func VMessage(message, name string, value string) {
	VPrint(fmt.Sprintf("%s %s", message, name), value)
}

func Found(name string, value string) {
	VMessage("found", name, value)
}

func NotFound(name string) {
	VMessage("did not find", name, "")
}

func CreateSpacerString(spaces int) (spacer string) {
	for i := 0; i < spaces; i++ {
		spacer += " "
	}
	return spacer
}

func Log(format string, a ...any) {
	fmt.Println(fmt.Sprintf(format, a...))
}
