package main

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/command"
	"cloud/pkg/config"
	"cloud/pkg/help"
	"cloud/pkg/util"
	"reflect"
)

var CommandMap = map[string]any{
	"init":    "Initialize",
	"destroy": "Destroy",
	"delete": map[string]any{
		"instance":  "DeleteInstances",
		"instances": "DeleteInstances",
		"key":       "DeleteKey",
		"keys":      "DeleteKey",
		"key-pair":  "DeleteKey",
		"key-pairs": "DeleteKey",
		"all": map[string]any{
			"keys":      "DeleteAllKeys",
			"key-pairs": "DeleteAllKeys",
		},
	},
	"list": map[string]any{
		"":          "PrintNodes",
		"instance":  "PrintNodes",
		"instances": "PrintNodes",
		"keys":      "ListKeys",
	},
	"create": map[string]any{
		"public": map[string]any{
			"instance": "CreatePublicInstance",
		},
		"private": map[string]any{
			"instance": "CreatePrivateInstance",
		},
		"key":      "CreateKeyPair",
		"key-pair": "CreateKeyPair",
	},
}

func main() {
	defer util.Recover()
	config.Load()
	amazon.InitEC2Client()
	traverse(CommandMap)
}

func traverse(commandMap map[string]any) {
	if val, ok := commandMap[args.PollOrEmpty()]; ok {
		if reflect.TypeOf(val).Kind() == reflect.String {
			exec(val)
		} else {
			traverse(val.(map[string]any))
		}
	} else {
		help.FatalHelpText()
	}
}

func exec(methodName any) {
	reflect.ValueOf(command.Commander{}).MethodByName(methodName.(string)).Call(nil)
}
