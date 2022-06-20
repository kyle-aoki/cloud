package args

import (
	"fmt"
	"strings"
)

type Flag struct {
	Key        string
	Abbr       string
	Value      string
	Standalone bool
}

var Flags = []Flag{
	{Key: "type", Abbr: "t", Value: ""},
	{Key: "gigs", Abbr: "g", Value: ""},
	{Key: "name", Abbr: "n", Value: ""},
	{Key: "ami", Abbr: "a", Value: ""},
	{Key: "script", Abbr: "c", Value: ""},
	{Key: "private", Abbr: "p", Value: "", Standalone: true},
}

func FlagValueOrDefault(flagKey string, defaultValue string) string {
	for i := range Flags {
		if Flags[i].Key == flagKey {
			if Flags[i].Value != "" {
				return Flags[i].Value
			}
			return defaultValue
		}
	}
	panic("flag not found: " + flagKey)
}

const FlagValueChars = "abcdefghijklmnopqrstuvwxyz0123456789"

func Contains(l []byte, e byte) bool {
	for i := range l {
		if l[i] == e {
			return true
		}
	}
	return false
}

func ContainsOnlyValidCharsOrPanic(s string) {
	for i := range s {
		if !Contains([]byte(FlagValueChars), s[i]) {
			panic("invalid characters: " + s)
		}
	}
}

func ParseFlags() {
FlagLoop:
	for i, flag := range Flags {
	ArgLoop:
		for j, argument := range args {

			standaloneFlag := fmt.Sprintf("--%s", flag.Key)
			if flag.Standalone && strings.Contains(argument, standaloneFlag) {
				flag.Value = "true"
				continue FlagLoop
			}

			abbrStandaloneFlag := fmt.Sprintf("-%s", flag.Abbr)
			if flag.Standalone && strings.Contains(argument, abbrStandaloneFlag) {
				flag.Value = "true"
				continue FlagLoop
			}

			fullFlagPrefix := fmt.Sprintf("--%s=", flag.Key)
			if strings.Contains(argument, fullFlagPrefix) {
				s := strings.Split(argument, "=")
				if len(s) != 2 {
					panic("invalid flag: " + argument)
				}
				ContainsOnlyValidCharsOrPanic(s[1])
				Flags[i].Value = s[1]
				continue FlagLoop
			}

			abbrFlagPrefix := fmt.Sprintf("-%s", flag.Abbr)
			if argument == abbrFlagPrefix {
				if j == len(args)-1 {
					panic("missing value for flag: " + abbrFlagPrefix)
				}
				v := args[j+1]
				ContainsOnlyValidCharsOrPanic(v)
				Flags[i].Value = v
				continue FlagLoop
			}

			continue ArgLoop
		}
	}
}
