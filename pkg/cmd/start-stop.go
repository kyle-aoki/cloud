package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
)

func StartInstance() {
	instNames := args.Collect()
	co := resource.New()
	for _, name := range instNames {
		co.StartInstance(name)
	}
}

func StopInstance() {
	instNames := args.Collect()
	co := resource.New()
	for _, name := range instNames {
		co.StopInstance(name)
	}
}
