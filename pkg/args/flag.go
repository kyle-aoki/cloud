package args

import (
	"strings"
)

type CharFlag struct {
	Name  string
	Value string
}

func ParseFlags(fullArgs []string) ([]CharFlag, []int) {
	var flags []CharFlag
	var args []int
	for i := 0; i < len(fullArgs); i++ {
		if isFlag(fullArgs[i]) {
			if isSingleDashFlag(fullArgs[i]) {
				if i == len(fullArgs)-1 {
					cf := CharFlag{Name: fullArgs[i][1:]}
					flags = append(flags, cf)
					break
				} else {
					cf := CharFlag{Name: fullArgs[i][1:], Value: fullArgs[i+1]}
					flags = append(flags, cf)
					i++
					continue
				}
			} else {
				cf := splitDoubleDashFlag(fullArgs[i])
				flags = append(flags, cf)
				continue
			}
		} else {
			args = append(args, i)
		}
	}
	return flags, args
}

func isFlag(s string) bool {
	if len(s) < 2 {
		return false
	}
	if s[0] == '-' {
		return true
	}
	return false
}

func isSingleDashFlag(s string) bool {
	if len(s) >= 2 {
		if s[0] == '-' && s[1] != '-' {
			return true
		}
	}
	return false
}

func splitDoubleDashFlag(s string) CharFlag {
	parts := strings.Split(s, "=")
	if len(parts) < 2 {
		panic("invalid flag: " + s)
	}
	return CharFlag{Name: parts[0][2:], Value: parts[1]}
}
