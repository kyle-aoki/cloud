package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"io/ioutil"
	"log"
)

func CreateInstance() {
	log.Println("creating instance...")
	co := resource.NewCloudOperator()

	defaultName := args.IsEmpty(*args.Flags.Name, co.NextInstanceName())
	util.Log("using instance name: %s", defaultName)

	startUpScript := ReadScriptFile(*args.Flags.Script)
	util.Log("using script file: %s", *args.Flags.Script)

	rii := &resource.RunInstanceInput{
		Name:             defaultName,
		SubnetId:         co.SubnetId(*args.Flags.Private || *args.Flags.P),
		InstanceType:     *args.Flags.InstType,
		Size:             util.StringToInt(*args.Flags.Gigs),
		SecurityGroupIds: []*string{co.SecurityGroupOrPanic("22").GroupId},
		UserData:         startUpScript,
	}
	util.Log("create instance input: %v", *rii)

	_ = resource.RunInstance(rii)
	fmt.Println(defaultName)
}

func ReadScriptFile(path string) string {
	if path == "" {
		return ""
	}
	bytes, err := ioutil.ReadFile(path)
	util.MustExec(err)
	return string(bytes)
}
