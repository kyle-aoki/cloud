package util

import (
	"fmt"
)

const PleaseVerify = "Please verify config at ~/.cloud is correct."

func PanicVerify(format string, a ...any) {
	panic(fmt.Sprintf("%v\n%v", fmt.Sprintf(format, a...), PleaseVerify))
}
