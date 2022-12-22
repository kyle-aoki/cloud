package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func OpenPort(
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
	util.Check(err)
}

func ClosePort(
	instance *ec2.Instance,
	port string,
) {
	var newSecurityGroups []*string
	for _, groupIdentifier := range instance.SecurityGroups {
		if groupIdentifier.GroupName != nil && *groupIdentifier.GroupName == port {
			continue
		}
		newSecurityGroups = append(newSecurityGroups, groupIdentifier.GroupId)
	}
	_, err := amazon.EC2().ModifyInstanceAttribute(&ec2.ModifyInstanceAttributeInput{
		InstanceId: instance.InstanceId,
		Groups:     newSecurityGroups,
	})
	util.Check(err)
}

func SecurityGroupByNameOrPanic(sgs []*ec2.SecurityGroup, groupName string) *ec2.SecurityGroup {
	util.Log("searching for security group: %s", groupName)
	for _, sg := range sgs {
		if sg.GroupName != nil && *sg.GroupName == groupName {
			return sg
		}
	}
	panic(fmt.Sprintf("failed to find security group %s", groupName))
}

func GetSecurityGroupIdByNameOrNil(lr *LabResources, name string) *ec2.SecurityGroup {
	for _, sg := range lr.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	return nil
}

func SecurityGroupExists(sgs []*ec2.SecurityGroup, name string) bool {
	for _, sg := range sgs {
		if sg.GroupName != nil && *sg.GroupName == name {
			return true
		}
	}
	return false
}

func ValidatePort(portString string) (portInt int) {
	portInt64, err := strconv.ParseInt(portString, 10, 32)
	if err != nil {
		panic("invalid port")
	}
	portInt = int(portInt64)
	if portInt > 65535 || portInt < 1 {
		panic("invalid port")
	}
	return portInt
}
