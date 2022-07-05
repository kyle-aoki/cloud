package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"io/ioutil"
	"log"
)

func Run() {
	log.Println("creating instance...")
	lr := resource.NewLabResources()
	lr.Instances = resource.FindInstances()
	lr.PublicSubnet = resource.FindSubnet(resource.CloudLabPublicSubnet)
	lr.PrivateSubnet = resource.FindSubnet(resource.CloudLabPrivateSubnet)
	lr.SecurityGroups = resource.FindAllSecurityGroups()

	var name string
	if args.FlagExists(args.Flags.Name) {
		if resource.NameExists(lr.Instances, *args.Flags.Name) {
			panic("name taken")
		}
		name = *args.Flags.Name
	} else {
		name = resource.NextInstanceName(lr.Instances)
	}
	util.Log("using instance name: %s", name)

	startUpScript := ReadScriptFile(*args.Flags.Script)
	util.Log("using script file: %s", *args.Flags.Script)

	port22 := resource.SecurityGroupByNameOrPanic(lr.SecurityGroups, "22").GroupId

	rii := &resource.RunInstanceInput{
		Name:             name,
		SubnetId:         resource.SelectSubnet(lr, *args.Flags.Private || *args.Flags.P),
		InstanceType:     *args.Flags.InstType,
		Size:             util.StringToInt(*args.Flags.Gigs),
		SecurityGroupIds: []*string{port22},
		UserData:         startUpScript,
	}
	util.Log("create instance input: %v", *rii)

	_ = resource.RunInstance(rii)
	fmt.Println(name)
}

func ReadScriptFile(path string) string {
	if path == "" {
		return ""
	}
	bytes, err := ioutil.ReadFile(path)
	util.MustExec(err)
	return string(bytes)
}
