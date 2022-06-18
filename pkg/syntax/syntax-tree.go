package syntax

import (
	cmd "cloudlab/pkg/command"
)

type Cmd struct {
	fn          func()
	args        string
	order       int
	explanation string
	fullCommand string
}

var SyntaxTree = map[string]any{
	"version": Cmd{fn: cmd.Version, order: 11},
	"info":    Cmd{fn: cmd.Info, order: 12, args: "discover existing cloudlab resources"},
	"init":    Cmd{fn: cmd.InitializeCloudLabResources, order: 13, args: "initialize cloudlab resources"},
	"destroy": Cmd{fn: cmd.DestroyCloudLabResources, order: 14, args: "delete all cloudlab resources"},

	"list": map[string]any{
		"":      Cmd{fn: cmd.ListInstances, order: 21},
		"nodes": Cmd{fn: cmd.ListInstances, order: 22},
		"keys":  Cmd{fn: cmd.ListKeys, order: 23},
	},

	"create": map[string]any{
		"public": map[string]any{
			"node": Cmd{fn: cmd.CreatePublicInstance, order: 31},
		},
		"private": map[string]any{
			"node": Cmd{fn: cmd.CreatePrivateInstance, order: 32},
		},
		"key": Cmd{fn: cmd.CreateKeyPair, order: 33},
	},

	"delete": map[string]any{
		"node": Cmd{fn: cmd.DeleteInstances, args: "<nodes>...", order: 40},
		"key":  Cmd{fn: cmd.DeleteKey, args: "<keys>...", order: 41},
		"all": map[string]any{
			"keys":  Cmd{fn: cmd.DeleteAllKeys, order: 42},
			"nodes": Cmd{fn: cmd.DeleteAllInstances, order: 43},
		},
	},

	"open": map[string]any{
		"port": Cmd{fn: cmd.OpenPort, args: "<port> <node>", order: 51},
	},

	"close": map[string]any{
		"port": Cmd{fn: cmd.ClosePort, args: "<port> <node>", order: 52},
	},
}
