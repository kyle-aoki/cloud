package syntax

import (
	"cloud/pkg/command"
)

type Command struct {
	fn          func()
	args        string
	order       int
	explanation string
	fullCommand string
}

var SyntaxTree = map[string]any{
	"version": Command{fn: command.Version, order: 11},
	"info":    Command{fn: command.Info, order: 12, args: "discover existing cloudlab resources"},
	"init":    Command{fn: command.InitializeCloudLabResources, order: 13, args: "initialize cloudlab resources"},
	"destroy": Command{fn: command.DestroyCloudLabResources, order: 14, args: "delete all cloudlab resources"},

	"list": map[string]any{
		"":      Command{fn: command.ListInstances, order: 21},
		"nodes": Command{fn: command.ListInstances, order: 22},
		"keys":  Command{fn: command.ListKeys, order: 23},
	},

	"create": map[string]any{
		"public": map[string]any{
			"node": Command{fn: command.CreatePublicInstance, order: 31},
		},
		"private": map[string]any{
			"node": Command{fn: command.CreatePrivateInstance, order: 32},
		},
		"key": Command{fn: command.CreateKeyPair, order: 33},
	},

	"delete": map[string]any{
		"node": Command{fn: command.DeleteInstances, args: "<nodes>...", order: 40},
		"key":  Command{fn: command.DeleteKey, args: "<keys>...", order: 41},
		"all": map[string]any{
			"keys":  Command{fn: command.DeleteAllKeys, order: 42},
			"nodes": Command{fn: command.DeleteAllInstances, order: 43},
		},
	},

	"open": map[string]any{
		"port": Command{fn: command.OpenPort, args: "<port> <node>", order: 51},
	},

	"close": map[string]any{
		"port": Command{fn: command.ClosePort, args: "<port> <node>", order: 52},
	},
}
