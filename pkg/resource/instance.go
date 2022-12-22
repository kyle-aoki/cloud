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
}

func RunInstance(rii *RunInstanceInput) *ec2.Instance {
	rio, err := amazon.EC2().RunInstances(&ec2.RunInstancesInput{
		SubnetId: &rii.SubnetId,
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
			DeviceName: util.Ptr("/dev/sda1"),
			Ebs:        &ec2.EbsBlockDevice{VolumeSize: util.Ptr(int64(rii.Size))}},
		},
		ImageId:          util.Ptr(amazon.Ubuntu2204Ami()),
		MinCount:         util.Ptr(int64(1)),
		MaxCount:         util.Ptr(int64(1)),
		KeyName:          util.Ptr(CloudLabKeyPair),
		SecurityGroupIds: rii.SecurityGroupIds,
		InstanceType:     &rii.InstanceType,
		TagSpecifications: CreateTagSpecs("instance", map[string]string{
			"Name":                   rii.Name,
			IsCloudLabInstanceTagKey: IsCloudLabInstanceTagVal,
		}),
	})
	util.Check(err)
	return rio.Instances[0]
}

// i1, i2, i3...
func NextInstanceName(instances []*ec2.Instance) string {
	var max int64
	for _, inst := range instances {
		n := FindNameTagValue(inst.Tags)
		if n == nil {
			continue
		}
		num := (*n)[1:]
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

func StartInstance(name string) {
	inst := FindInstanceByName(name)
	if inst == nil {
		panic("instance not found")
	}
	_, err := amazon.EC2().StartInstances(&ec2.StartInstancesInput{
		InstanceIds: []*string{inst.InstanceId},
	})
	util.Check(err)
	fmt.Println(name)
}

func StopInstance(name string) {
	inst := FindInstanceByName(name)
	if inst == nil {
		panic("instance not found")
	}
	_, err := amazon.EC2().StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{inst.InstanceId},
	})
	util.Check(err)
	fmt.Println(name)
}

func TerminateInstances(instances []*ec2.Instance) {
	var instanceIds []*string
	for i := 0; i < len(instances); i++ {
		instanceIds = append(instanceIds, instances[i].InstanceId)
	}
	_, err := amazon.EC2().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	})
	util.Check(err)
}

func NameExists(instances []*ec2.Instance, name string) bool {
	for _, inst := range instances {
		if NameTagEquals(inst.Tags, name) {
			return true
		}
	}
	return false
}

func HasPortOpen(instance *ec2.Instance, port string) bool {
	for _, sg := range instance.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == port {
			return true
		}
	}
	return false
}

func InPrivateSubnet(instance *ec2.Instance, lr *LabResources) bool {
	return *lr.PrivateSubnet.SubnetId == *instance.SubnetId
}
