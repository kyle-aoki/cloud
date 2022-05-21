package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/defaults"
	"cloud/pkg/util"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"
)

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

func OpenPort() {
	port := args.Poll()
	instnaceName := args.Poll()

	cldo := defaults.Start()

	sg := cldo.GetSecurityGroupIdByNameOrNil(port)

	if sg == nil {
		portInt := ValidatePort(port)
		cldo.CreateSecurityGroup(port, portInt)
	}

	instance := cldo.GetInstanceByName(instnaceName)
	securityGroup := cldo.GetSecurityGroupIdByNameOrPanic(port)

	cldo.AssignSecurityGroup(instance, securityGroup)
	fmt.Println(fmt.Sprintf("opened port %s on node %s", port, instnaceName))
}

func ClosePort() {
	cldo := defaults.Start()

	port := args.Poll()
	instanceName := args.Poll()

	_ = ValidatePort(port)

	var instance *ec2.Instance = cldo.FindInstanceByName(instanceName)

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
	util.MustExec(err)

	fmt.Println(fmt.Sprintf("closed port %s on node %s", port, instanceName))
}
