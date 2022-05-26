package resource

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type CreateInstanceInput struct {
	Name               string
	Size               int
	SubnetId           string
	CurrentKeyPairName string
	SecurityGroupIds   []*string
	UserData           string
}

func CreateInstance(cii *CreateInstanceInput) *ec2.Instance {
	rio, err := amazon.EC2().RunInstances(&ec2.RunInstancesInput{
		SubnetId: util.StrPtr(cii.SubnetId),
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
			DeviceName: util.StrPtr("/dev/sda1"),
			Ebs:        &ec2.EbsBlockDevice{VolumeSize: util.IntToInt64Ptr(cii.Size)}},
		},
		ImageId:          util.StrPtr(amazon.GetAmi()),
		MinCount:         util.IntToInt64Ptr(1),
		MaxCount:         util.IntToInt64Ptr(1),
		KeyName:          util.StrPtr(cii.CurrentKeyPairName),
		SecurityGroupIds: cii.SecurityGroupIds,
		InstanceType:     util.StrPtr("t2.micro"),
		TagSpecifications: CreateTagSpecs("instance", map[string]string{
			"Name": cii.Name,
		}),
		UserData: util.StrPtr(cii.UserData),
	})
	util.MustExec(err)
	return rio.Instances[0]
}

func AssignSecurityGroup(
	instance *ec2.Instance,
	securityGroup *ec2.SecurityGroup,
) {
	var groupIds []*string
	for _, sgs := range instance.SecurityGroups {
		groupIds = append(groupIds, sgs.GroupId)
	}
	groupIds = append(groupIds, securityGroup.GroupId)
	_, err := amazon.EC2().ModifyInstanceAttribute(&ec2.ModifyInstanceAttributeInput{
		InstanceId: instance.InstanceId,
		Groups:     groupIds,
	})
	util.MustExec(err)
}
