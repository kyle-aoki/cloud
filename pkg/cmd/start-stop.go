package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
)

func StartInstance() {
	instNames := args.CollectOrPanic()
	for _, name := range instNames {
		resource.StartInstance(name)
	}
}

func StopInstance() {
	instNames := args.CollectOrPanic()
	for _, name := range instNames {
		resource.StopInstance(name)
	}
}
