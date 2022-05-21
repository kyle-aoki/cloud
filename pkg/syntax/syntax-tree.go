package syntax

import (
	"cloud/pkg/command"
)

type Command struct {
	fn   func()
	args string
}

var SyntaxTree = map[string]any{
	"ssh":     Command{fn: command.SSH},
	"init":    Command{fn: command.InitializeCloudLabResources},
	"destroy": Command{fn: command.DestroyCloudLabResources},
	"delete": map[string]any{
		"node": Command{fn: command.DeleteInstances, args: "<nodes>..."},
		"key":  Command{fn: command.DeleteKey, args: "<keys>..."},
		"all": map[string]any{
			"keys": Command{fn: command.DeleteAllKeys},
		},
	},
	"list": map[string]any{
		"":      Command{fn: command.ListInstances},
		"nodes": Command{fn: command.ListInstances},
		"keys":  Command{fn: command.ListKeys},
	},
	"create": map[string]any{
		"public": map[string]any{
			"node": Command{fn: command.CreatePublicInstance},
		},
		"private": map[string]any{
			"node": Command{fn: command.CreatePrivateInstance},
		},
		"key": Command{fn: command.CreateKeyPair},
	},
	"open": map[string]any{
		"port": Command{fn: command.OpenPort, args: "<port> <node>"},
	},
	"close": map[string]any{
		"port": Command{fn: command.ClosePort, args: "<port> <node>"},
	},
}
