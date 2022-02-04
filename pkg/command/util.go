package command

import "fmt"

var f func(format string, a ...any) string

func init() {
	f = fmt.Sprintf
}
