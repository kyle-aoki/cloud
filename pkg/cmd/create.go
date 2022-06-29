package cmd

func CreateInstance() {
	// co := resource.NewCloudOperator()

	// name := args.FlagValueOrDefault("name", co.NextInstanceName())
	// instType := args.FlagValueOrDefault("type", "t2.nano")
	// ami := args.FlagValueOrDefault("ami", amazon.UbuntuAmi())
	// gigs := args.FlagValueOrDefault("gigs", "8")
	// scriptPath := args.FlagValueOrDefault("script", "")
	// private := args.FlagValueOrDefault("private", "false")

	// startUpScript := ReadScriptFile(scriptPath)

	// cii := &resource.CreateInstanceInput{
	// 	TagSpecifications: resource.CreateTagSpecs("instance", map[string]string{
	// 		"Name":                            name,
	// 		resource.IsCloudLabInstanceTagKey: resource.IsCloudLabInstanceTagVal,
	// 	}),
	// 	SubnetId:         co.SelectPrivateSubnet(private),
	// 	InstanceType:     instType,
	// 	Ami:              ami,
	// 	DeviceName:       "/dev/sda1",
	// 	MinCount:         1,
	// 	MaxCount:         1,
	// 	Size:             util.StringToInt(gigs),
	// 	KeyName:          resource.CloudLabKeyPair,
	// 	SecurityGroupIds: []*string{co.GetSecurityGroupIdByNameOrPanic("22").GroupId},
	// 	UserData:         startUpScript,
	// }

	// instance := resource.ExecuteCreateInstanceRequest(cii)
	// fmt.Println(*instance.InstanceId)
}

// func ReadScriptFile(path string) string {
// 	if path == "" {
// 		return ""
// 	}
// 	bytes, err := ioutil.ReadFile(path)
// 	util.MustExec(err)
// 	return string(bytes)
// }
