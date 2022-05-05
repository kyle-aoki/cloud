package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/config"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func Create() {
	configType := args.Poll()
	instanceNames := args.Collect()
	client := amazon.EC2()

	nc := config.Vars.Find(configType)

	for _, instanceName := range instanceNames {
		rio, err := client.RunInstances(&ec2.RunInstancesInput{
			BlockDeviceMappings: []*ec2.BlockDeviceMapping{
				{
					DeviceName: nc.DefaultDeviceName(),
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: nc.StorageSizeToInt64(),
					}},
			},
			ImageId:          &nc.AMI,
			MinCount:         util.IntToInt64Ptr(1),
			MaxCount:         util.IntToInt64Ptr(1),
			KeyName:          &nc.KeyPair,
			SecurityGroupIds: nc.SecurityGroupIds(),
			InstanceType:     &nc.InstanceType,
			SubnetId:         nc.SubnetId(),
			TagSpecifications: []*ec2.TagSpecification{
				{ResourceType: util.StrPtr("instance"), Tags: []*ec2.Tag{
					{Key: util.StrPtr("Name"), Value: util.StrPtr(instanceName)},
				}},
			},
			PrivateIpAddress: nc.GetPrivateIp(),
			UserData: nc.GetUserData(),
		})
		util.MustExec(err)

		fmt.Println(*rio.Instances[0].InstanceId)
	}
}
