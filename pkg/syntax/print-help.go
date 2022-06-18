package syntax

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var commands []Cmd
var spacing = 40

func HelpText() {
	traverseAndAppend(SyntaxTree, "lab")
	sortCommands()
	printCommands()
	os.Exit(1)
}

func sortCommands() {
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].order < commands[j].order
	})
}

func printCommands() {
	fmt.Println()
	fmt.Println(createSpacedString("command", "arguments/explanation"))
	fmt.Println(strings.Repeat("-", 60))
	for i := range commands {
		fmt.Println(commands[i].fullCommand)
	}
	fmt.Println()
}

// recursive
func traverseAndAppend(syntaxTree map[string]any, prevkey string) {
	for key, value := range syntaxTree {
		switch value.(type) {
		case Cmd:
			formLine(concatKey(prevkey, key), value.(Cmd))
		default:
			traverseAndAppend(value.(map[string]any), concatKey(prevkey, key))
		}
	}
}

func createSpacedString(left string, right string) string {
	spaces := spacing - len(left)
	spaceString := strings.Repeat(" ", spaces)
	return fmt.Sprintf("%s%s%s", left, spaceString, right)
}

func formLine(left string, command Cmd) {
	command.fullCommand = createSpacedString(left, command.args)
	commands = append(commands, command)
}

var lineBreak = func() { fmt.Println() }
var concatKey = func(prev string, key string) string { return fmt.Sprintf("%s %s", prev, key) }
