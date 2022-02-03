package util

import (
	"cloud/pkg/errs"
	"fmt"
)

func PanicVerify(format string, a ...any) {
	panic(fmt.Sprintf("%v\n%v", fmt.Sprintf(format, a...), errs.PleaseVerify))
}
