package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/ami"
	"cloud/pkg/defaults"
	"cloud/pkg/names"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func Create() {
	cldo := &defaults.CloudLabDefaultsOperator{}
	cldo.FindAll()
	cldo.FindAllCloudLabKeyPairs()

	rio, err := amazon.EC2().RunInstances(&ec2.RunInstancesInput{
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
			DeviceName: util.StrPtr("/dev/sda1"),
			Ebs:        &ec2.EbsBlockDevice{VolumeSize: util.IntToInt64Ptr(8)}},
		},
		ImageId:  util.StrPtr(ami.GetAmi()),
		MinCount: util.IntToInt64Ptr(1),
		MaxCount: util.IntToInt64Ptr(1),
		KeyName:  util.StrPtr(cldo.GetCurrentCloudLabKeyPairName()),
		SecurityGroupIds: []*string{
			util.StrPtr(*cldo.GetSecurityGroupIdByName("allow-port-22").GroupId),
		},
		InstanceType:      util.StrPtr("t2.micro"),
		SubnetId:          cldo.PublicSubnet.SubnetId,
		TagSpecifications: defaults.CreateNameTagSpec("instance", names.GetRandomName()),
	})
	util.MustExec(err)
	fmt.Println("created instance " + *rio.Instances[0].InstanceId)
}
