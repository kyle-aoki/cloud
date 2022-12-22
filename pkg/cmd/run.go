package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"log"
)

func Run() {
	log.Println("creating instance...")
	lr := resource.NewLabResources()
	lr.Instances = resource.FindInstances()
	lr.PublicSubnet = resource.FindPublicSubnet()
	lr.PrivateSubnet = resource.FindPrivateSubnet()
	lr.SecurityGroups = resource.FindAllSecurityGroups()

	var name string
	if args.Flags.InstanceName != nil && *args.Flags.InstanceName != "" {
		if resource.NameExists(lr.Instances, *args.Flags.InstanceName) {
			panic(fmt.Sprintf("cannot name instance '%s': name already taken", *args.Flags.InstanceName))
		}
		name = *args.Flags.InstanceName
	} else {
		name = resource.NextInstanceName(lr.Instances)
	}
	util.Log("using instance name: %s", name)

	port22 := resource.SecurityGroupByNameOrPanic(lr.SecurityGroups, "22").GroupId

	rii := &resource.RunInstanceInput{
		Name:             name,
		SubnetId:         resource.SelectSubnet(lr, args.Flags.Private),
		InstanceType:     args.Flags.InstanceType,
		Size:             util.StringToInt(args.Flags.Gigabytes),
		SecurityGroupIds: []*string{port22},
	}
	util.Log("create instance input: %v", *rii)

	_ = resource.RunInstance(rii)

	hostConfig := addInstanceToSshConfig(name, lr)

	if args.Flags.Quiet {
		fmt.Println(name)
	} else {
		fmt.Println(hostConfig)
	}
}
