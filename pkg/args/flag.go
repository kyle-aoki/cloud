package args

import (
	"strings"
)

type CharFlag struct {
	Flag  string
	Value string
}

func Par(fullArgs []string) ([]CharFlag, []int) {
	var flags []CharFlag
	var args []int
	for i := 0; i < len(fullArgs); i++ {
		if IsFlag(fullArgs[i]) {
			if IsSingleDashFlag(fullArgs[i]) {
				if i == len(fullArgs)-1 {
					cf := CharFlag{Flag: fullArgs[i][1:]}
					flags = append(flags, cf)
					break
				} else {
					cf := CharFlag{Flag: fullArgs[i][1:], Value: fullArgs[i+1]}
					flags = append(flags, cf)
					i++
					continue
				}
			} else {
				cf := SplitDoubleDashFlag(fullArgs[i])
				flags = append(flags, cf)
				continue
			}
		} else {
			args = append(args, i)
		}
	}
	return flags, args
}

func IsFlag(s string) bool {
	if len(s) < 2 {
		return false
	}
	if s[0] == '-' {
		return true
	}
	return false
}

func IsSingleDashFlag(s string) bool {
	if len(s) >= 2 {
		if s[0] == '-' && s[1] != '-' {
			return true
		}
	}
	return false
}

func SplitDoubleDashFlag(s string) CharFlag {
	parts := strings.Split(s, "=")
	if len(parts) < 2 {
		panic("invalid flag: " + s)
	}
	return CharFlag{Flag: parts[0][2:], Value: parts[1]}
}
