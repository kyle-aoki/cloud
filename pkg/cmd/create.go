package cmd

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"io/ioutil"
)

func CreateInstance() {
	co := resource.NewCloudOperator()

	defaultName := co.NextInstanceName()

	ami := args.StrFlag(amazon.UbuntuAmi(), "ami", "a")
	instType := args.StrFlag("t2.nano", "type", "t")
	isPrivateSubnet := args.BoolFlag("private", "p")
	name := args.StrFlag(defaultName, "name", "n")
	scriptPath := args.StrFlag("", "script", "s")
	gigs := args.StrFlag("8", "gigs", "g")

	startUpScript := ReadScriptFile(scriptPath)

	cii := &resource.CreateInstanceInput{
		TagSpecifications: resource.CreateTagSpecs("instance", map[string]string{
			"Name":                            name,
			resource.IsCloudLabInstanceTagKey: resource.IsCloudLabInstanceTagVal,
		}),
		SubnetId:         co.UsePrivateSubnet(isPrivateSubnet),
		InstanceType:     instType,
		Ami:              ami,
		DeviceName:       "/dev/sda1",
		MinCount:         1,
		MaxCount:         1,
		Size:             util.StringToInt(gigs),
		KeyName:          resource.CloudLabKeyPair,
		SecurityGroupIds: []*string{co.GetSecurityGroupIdByNameOrPanic("22").GroupId},
		UserData:         startUpScript,
	}

	_ = resource.ExecuteCreateInstanceRequest(cii)
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
