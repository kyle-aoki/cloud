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
	co := resource.New()
	co.Rs.Instances = co.Finder.FindInstances()
	co.Rs.SecurityGroups = co.Finder.FindAllSecurityGroups()

	var name string
	if args.FlagExists(args.Flags.Name) {
		if resource.NameExists(co.Rs.Instances, *args.Flags.Name) {
			panic("name taken")
		}
		name = *args.Flags.Name
	} else {
		name = resource.NextInstanceName(co.Rs.Instances)
	}
	util.Log("using instance name: %s", name)

	startUpScript := ReadScriptFile(*args.Flags.Script)
	util.Log("using script file: %s", *args.Flags.Script)

	port22 := resource.SecurityGroupByNameOrPanic(co.Rs.SecurityGroups, "22").GroupId

	rii := &resource.RunInstanceInput{
		Name:             name,
		SubnetId:         co.SubnetId(*args.Flags.Private || *args.Flags.P),
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
