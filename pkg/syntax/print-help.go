package syntax

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var lines []string

func helpText() {
	lineBreak()
	formLine("command", "arguments")
	lines = append(lines, strings.Repeat("-", 60))
	traverseAndAppend(SyntaxTree, "lab")
	printLines()
	lineBreak()
	os.Exit(1)
}

func printLines() {
	firstTwoLines := lines[:2]
	remLines := lines[2:]
	sort.Strings(remLines)
	firstTwoLines = append(firstTwoLines, remLines...)
	for i := range firstTwoLines {
		fmt.Println(firstTwoLines[i])
	}
}

// recursive
func traverseAndAppend(syntaxTree map[string]any, prevkey string) {
	for key, value := range syntaxTree {
		switch value.(type) {
		case Command:
			formLine(concatKey(prevkey, key), value.(Command).args)
		default:
			traverseAndAppend(value.(map[string]any), concatKey(prevkey, key))
		}
	}
}

func formLine(left string, right string) {
	spaces := 40 - len(left)
	spaceString := strings.Repeat(" ", spaces)
	lines = append(lines, fmt.Sprintf("%s%s%s", left, spaceString, right))
}

var lineBreak = func() { fmt.Println() }
var concatKey = func(prev string, key string) string { return fmt.Sprintf("%s %s", prev, key) }
