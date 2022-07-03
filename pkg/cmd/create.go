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

	defaultName := args.IsEmpty(*args.Flags.Name, co.NextInstanceName())

	startUpScript := ReadScriptFile(*args.Flags.Script)

	cii := &resource.CreateInstanceInput{
		TagSpecifications: resource.CreateTagSpecs("instance", map[string]string{
			"Name":                            defaultName,
			resource.IsCloudLabInstanceTagKey: resource.IsCloudLabInstanceTagVal,
		}),
		SubnetId:         co.UsePrivateSubnet(*args.Flags.Private || *args.Flags.P),
		InstanceType:     *args.Flags.InstType,
		Ami:              amazon.UbuntuAmi(),
		DeviceName:       "/dev/sda1",
		MinCount:         1,
		MaxCount:         1,
		Size:             util.StringToInt(*args.Flags.Gigs),
		KeyName:          resource.CloudLabKeyPair,
		SecurityGroupIds: []*string{co.GetSecurityGroupIdByNameOrPanic("22").GroupId},
		UserData:         startUpScript,
	}

	_ = resource.ExecuteCreateInstanceRequest(cii)
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
