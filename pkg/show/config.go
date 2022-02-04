package show

import (
	"cloud/pkg/config"
	"cloud/pkg/tab"
	"fmt"
)

func ShowConfig() {
	tab.Print("Variable Name\tValue")
	for _, subnet := range config.Vars.Variables.Subnets {
		tab.Print(format(subnet.Name, subnet.Value))
	}
	for _, sg := range config.Vars.Variables.SecurityGroups {
		tab.Print(format(sg.Name, sg.Value))
	}
	for _, ami := range config.Vars.Variables.AMIs {
		tab.Print(format(ami.Name, ami.Value))
	}
	for _, nodeConfig := range config.Vars.NodeConfigs {
		tab.Print(fmt.Sprintf("\n%v:\t%v\n%v:\t%v\n%v:\t%v\n%v:\t%v\n%v:\t%v\n%v:\t%v\n%v:\t%v\n",
			"Name", nodeConfig.Name,
			"Subnet", nodeConfig.Subnet,
			"SecurityGroups", nodeConfig.SecurityGroups,
			"KeyPair", nodeConfig.KeyPair,
			"InstanceType", nodeConfig.InstanceType,
			"StorageSize", nodeConfig.StorageSize,
			"Ami", nodeConfig.Ami,
		))
	}
	tab.Flush()
}

func format(name string, value string) string {
	return fmt.Sprintf("%v\t%v", name, value)
}
