package util

import "fmt"

// ########################################
const messageLength = 50

func VPrint(message string, value string) {
	if len(message) > messageLength {
		fmt.Println(fmt.Sprintf("%s %s", message, value))
		return
	}
	spaces := messageLength - len(message)
	fmt.Println(fmt.Sprintf("%s%s%s", message, CreateSpacerString(spaces), value))
}

func CreateSpacerString(spaces int) (spacer string) {
	for i := 0; i < spaces; i++ {
		spacer += " "
	}
	return spacer
}
