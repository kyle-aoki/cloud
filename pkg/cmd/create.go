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
	ro := resource.NewResourceOperator()

	name := args.FlagValueOrDefault("name", ro.NextInstanceName())
	instType := args.FlagValueOrDefault("type", "t2.nano")
	ami := args.FlagValueOrDefault("ami", amazon.UbuntuAmi())
	gigs := args.FlagValueOrDefault("gigs", "8")
	scriptPath := args.FlagValueOrDefault("script", "")
	private := args.FlagValueOrDefault("private", "false")

	startUpScript := ReadScriptFile(scriptPath)

	cii := &resource.CreateInstanceInput{
		TagSpecifications: resource.CreateTagSpecs("instance", map[string]string{
			"Name":                            name,
			resource.IsCloudLabInstanceTagKey: resource.IsCloudLabInstanceTagVal,
		}),
		SubnetId:         ro.SelectPrivateSubnet(private),
		InstanceType:     instType,
		Ami:              ami,
		DeviceName:       "/dev/sda1",
		MinCount:         1,
		MaxCount:         1,
		Size:             util.StringToInt(gigs),
		KeyName:          resource.CloudLabKeyPair,
		SecurityGroupIds: []*string{ro.GetSecurityGroupIdByNameOrPanic("22").GroupId},
		UserData:         startUpScript,
	}

	instance := resource.ExecuteCreateInstanceRequest(cii)
	fmt.Println(*instance.InstanceId)
}

func ReadScriptFile(path string) string {
	if path == "" {
		return ""
	}
	bytes, err := ioutil.ReadFile(path)
	util.MustExec(err)
	return string(bytes)
}
