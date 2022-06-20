package syntax

import (
	"cloudlab/pkg/cmd"
)

type Cmd struct {
	fn          func()
	args        []string
	order       int
	explanation string
	fullCommand string
}

var SyntaxTree = map[string]any{
	"version": Cmd{fn: cmd.Version, order: 11},
	"info":    Cmd{fn: cmd.Info, order: 12, args: []string{"discover existing cloudlab resources"}},
	"init":    Cmd{fn: cmd.InitializeCloudLabResources, order: 13, args: []string{"initialize cloudlab resources"}},
	"destroy": Cmd{fn: cmd.DestroyCloudLabResources, order: 14, args: []string{"delete all cloudlab resources"}},

	"list": Cmd{fn: cmd.ListInstances, order: 21, args: []string{"list instances"}},

	"run": Cmd{fn: cmd.CreateInstance, order: 31, args: []string{
		"all flags optional",
		"-n, --name=<string>",
		"-p, --private",
		"-g, --gigs=<storage>",
		"-t, --type=<t2.nano, t2.micro, etc>",
		"-a, --ami=<amazon-machine-image>",
		"-s, --script=<script-file>"},
	},

	"delete": Cmd{fn: cmd.DeleteInstances, args: []string{"<names>..."}, order: 40},

	"open": map[string]any{
		"port": Cmd{fn: cmd.OpenPort, args: []string{"<port> <name>"}, order: 51},
	},

	"close": map[string]any{
		"port": Cmd{fn: cmd.ClosePort, args: []string{"<port> <name>"}, order: 52},
	},
}
