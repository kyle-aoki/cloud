package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type RunInstanceInput struct {
	Name             string
	SubnetId         string
	InstanceType     string // t2.nano
	Size             int
	SecurityGroupIds []*string
	UserData         string
}

func RunInstance(rii *RunInstanceInput) *ec2.Instance {
	rio, err := amazon.EC2().RunInstances(&ec2.RunInstancesInput{
		SubnetId: util.StrPtr(rii.SubnetId),
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
			DeviceName: util.StrPtr("/dev/sda1"),
			Ebs:        &ec2.EbsBlockDevice{VolumeSize: util.IntToInt64Ptr(rii.Size)}},
		},
		ImageId:          util.StrPtr(amazon.UbuntuAmi()),
		MinCount:         util.IntToInt64Ptr(1),
		MaxCount:         util.IntToInt64Ptr(1),
		KeyName:          util.StrPtr(CloudLabKeyPair),
		SecurityGroupIds: rii.SecurityGroupIds,
		InstanceType:     util.StrPtr(rii.InstanceType),
		UserData:         util.StrPtr(rii.UserData),
		TagSpecifications: CreateTagSpecs("instance", map[string]string{
			"Name":                   rii.Name,
			IsCloudLabInstanceTagKey: IsCloudLabInstanceTagVal,
		}),
	})
	util.MustExec(err)
	return rio.Instances[0]
}

// i1, i2, i3 ...
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

func (co *AWSCloudOperator) StartInstance(name string) {
	inst := co.Finder.FindInstanceByName(name)
	if inst == nil {
		panic("instance not found")
	}
	_, err := amazon.EC2().StartInstances(&ec2.StartInstancesInput{
		InstanceIds: []*string{inst.InstanceId},
	})
	util.MustExec(err)
	fmt.Println("started " + name)
}

func (co *AWSCloudOperator) StopInstance(name string) {
	inst := co.Finder.FindInstanceByName(name)
	if inst == nil {
		panic("instance not found")
	}
	_, err := amazon.EC2().StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{inst.InstanceId},
	})
	util.MustExec(err)
	fmt.Println("stopped " + name)
}

func (co *AWSCloudOperator) TerminateInstance(id *string) {
	_, err := amazon.EC2().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{id},
	})
	util.MustExec(err)
}
