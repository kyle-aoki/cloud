package main

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/command"
	"cloud/pkg/util"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	fn   func()
	args string
}

var CommandMap = map[string]any{
	"init":    Command{fn: command.InitializeCloudLabResources, args: ""},
	"destroy": Command{fn: command.DestroyCloudLabResources, args: ""},
	"delete": map[string]any{
		"node":  Command{fn: command.DeleteInstances, args: "<instance-names>..."},
		"nodes": Command{fn: command.DeleteInstances, args: "<instance-names>..."},
		"key":       Command{fn: command.DeleteKey, args: "<key-names>..."},
		"keys":      Command{fn: command.DeleteKey, args: "<key-names>..."},
		"key-pair":  Command{fn: command.DeleteKey, args: "<key-names>..."},
		"key-pairs": Command{fn: command.DeleteKey, args: "<key-names>..."},
		"all": map[string]any{
			"keys":      Command{fn: command.DeleteAllKeys, args: ""},
			"key-pairs": Command{fn: command.DeleteAllKeys, args: ""},
		},
	},
	"list": map[string]any{
		"":          Command{fn: command.ListInstances, args: ""},
		"node":  Command{fn: command.ListInstances, args: ""},
		"nodes": Command{fn: command.ListInstances, args: ""},
		"keys":      Command{fn: command.ListKeys, args: ""},
		"key-pairs": Command{fn: command.ListKeys, args: ""},
	},
	"create": map[string]any{
		"public": map[string]any{
			"node": Command{fn: command.CreatePublicInstance, args: ""},
		},
		"private": map[string]any{
			"node": Command{fn: command.CreatePrivateInstance, args: ""},
		},
		"key":      Command{fn: command.CreateKeyPair, args: ""},
		"key-pair": Command{fn: command.CreateKeyPair, args: ""},
	},
	"open": map[string]any{
		"port": Command{fn: command.OpenPort, args: "<port> <node>"},
	},
}

func main() {
	defer util.Recover()
	amazon.InitEC2Client()

	traverse(CommandMap)
}

func HelpText() {
	fmt.Println()
	ArrowTabPrint("command", "arguments")
	fmt.Println(strings.Repeat("-", 60))
	PrintCommands(CommandMap, "cloudlab")
	fmt.Println()
	os.Exit(1)
}

func ConcatKey(prev string, key string) string {
	return fmt.Sprintf("%s %s", prev, key)
}

func ArrowTabPrint(left string, right string) {
	spaces := 40 - len(left)
	spaceString := strings.Repeat(" ", spaces)
	fmt.Println(fmt.Sprintf("%s%s%s", left, spaceString, right))
}

func PrintCommands(commandMap map[string]any, prevkey string) {
	for key, value := range commandMap {
		switch value.(type) {
		case Command:
			ArrowTabPrint(ConcatKey(prevkey, key), value.(Command).args)
		default:
			PrintCommands(value.(map[string]any), ConcatKey(prevkey, key))
		}
	}
}

func traverse(commandMap map[string]any) {
	if val, ok := commandMap[args.PollOrEmpty()]; ok {
		switch val.(type) {
		case Command:
			val.(Command).fn()
		default:
			traverse(val.(map[string]any))
		}
	} else {
		HelpText()
	}
}
