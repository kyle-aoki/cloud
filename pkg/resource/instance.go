package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type CreateInstanceInput struct {
	Name              string
	SubnetId          string
	InstanceType      string // t2.nano
	Ami               string
	DeviceName        string // Ubuntu: /dev/sda1
	MinCount          int
	MaxCount          int
	Size              int
	KeyName           string
	SecurityGroupIds  []*string
	UserData          string
	TagSpecifications []*ec2.TagSpecification
}

func ExecuteCreateInstanceRequest(cii *CreateInstanceInput) *ec2.Instance {
	rio, err := amazon.EC2().RunInstances(&ec2.RunInstancesInput{
		SubnetId: util.StrPtr(cii.SubnetId),
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
			DeviceName: util.StrPtr(cii.DeviceName),
			Ebs:        &ec2.EbsBlockDevice{VolumeSize: util.IntToInt64Ptr(cii.Size)}},
		},
		ImageId:           util.StrPtr(cii.Ami),
		MinCount:          util.IntToInt64Ptr(cii.MinCount),
		MaxCount:          util.IntToInt64Ptr(cii.MaxCount),
		KeyName:           util.StrPtr(cii.KeyName),
		SecurityGroupIds:  cii.SecurityGroupIds,
		InstanceType:      util.StrPtr(cii.InstanceType),
		TagSpecifications: cii.TagSpecifications,
		UserData:          util.StrPtr(cii.UserData),
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

func (co *AWSCloudOperator) NextInstanceName() string {
	var max int64
	for _, inst := range co.Rs.Instances {
		n := *FindNameTagValue(inst.Tags)
		num := n[1:]
		i, err := strconv.ParseInt(num, 10, 32)
		if err != nil {
			continue
		}
		if i > max {
			max = i
		}
	}
	return fmt.Sprintf("i%v", max+1)
}

func (co *AWSCloudOperator) SelectPrivateSubnet(isPrivate string) string {
	if isPrivate == "true" {
		return *co.Rs.PrivateSubnet.SubnetId
	}
	return *co.Rs.PublicSubnet.SubnetId
}
