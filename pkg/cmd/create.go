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

	name := args.StrFlag("name", defaultName)
	instType := args.StrFlag("type", "t2.nano")
	ami := args.StrFlag("ami", amazon.UbuntuAmi())
	gigs := args.StrFlag("gigs", "8")
	scriptPath := args.StrFlag("script", "")
	isPrivateSubnet := args.BoolFlag("private")

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
