package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
)

func SSH() {
	names := args.CollectOrEmpty()
	allInstances := len(names) == 0

	lr := resource.NewLabResources()
	lr.Instances = resource.FindNonTerminatedInstances()

	for _, inst := range lr.Instances {
		instName := resource.FindNameTagValue(inst.Tags)
		var ip string
		if *args.Flags.P || *args.Flags.Private {
			ip = *inst.PrivateIpAddress
		} else {
			ip = *inst.PublicIpAddress
		}
		if !allInstances && instName != nil && util.Contains(*instName, names) {
			printSSHCommand(ip)
		}
		if allInstances {
			printSSHCommand(ip)
		}
	}
}

func printSSHCommand(ip string) {
	fmt.Println(fmt.Sprintf("ssh -i %s ubuntu@%s", resource.KeyFilePath(), ip))
}
