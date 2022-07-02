package args

import (
	"strings"
)

type Flag struct {
	Name  string
	Value string
}

func ExtractFlagsFromArgs(fullArgs []string) (flags []Flag, argIndicies []int) {
	for i := 0; i < len(fullArgs); i++ {
		a := fullArgs[i]
		if !isFlag(a) {
			argIndicies = append(argIndicies, i)
			continue
		}

		f := Flag{}
		singleDash := hasOneDash(a)
		lastArg := isLastArg(i, fullArgs)

		switch {
		case singleDash && lastArg:
			f.Name = fullArgs[i][1:]
		case singleDash && !lastArg:
			f.Name = fullArgs[i][1:]
			f.Value = fullArgs[i+1]
			i++
		case !singleDash:
			f.Name, f.Value = splitDoubleDashFlag(fullArgs[i])
		default:
			panic("something went wrong parsing flags")
		}

		flags = append(flags, f)
	}
	return flags, argIndicies
}

func isLastArg(currentArgIndex int, fullArgs []string) bool {
	return currentArgIndex == len(fullArgs)-1
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

func hasOneDash(s string) bool {
	if len(s) >= 2 {
		if s[0] == '-' && s[1] != '-' {
			return true
		}
	}
	return false
}

func splitDoubleDashFlag(s string) (string, string) {
	parts := strings.Split(s, "=")
	if len(parts) < 2 {
		panic("invalid flag: " + s)
	}
	return parts[0][2:], parts[1]
}
