package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"os"
	"strconv"
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
	instanceName := args.Poll()

	util.Log("port %s", port)
	util.Log("instanceName %s", instanceName)

	co := resource.NewCloudOperator()

	sg := co.GetSecurityGroupIdByNameOrNil(port)

	if sg == nil {
		portInt := ValidatePort(port)
		co.Creator.CreateSecurityGroup(co.Rs.Vpc, port, portInt)
		co.Rs.SecurityGroups = co.Finder.FindSecurityGroups()
	}

	instance := co.Finder.FindInstanceByName(instanceName)

	for _, sg := range instance.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == port {
			fmt.Printf("port %s already open on instance %s\n", port, instanceName)
			os.Exit(0)
		}
	}

	securityGroup := co.SecurityGroupOrPanic(port)

	resource.AssignSecurityGroup(instance, securityGroup)
	fmt.Println(fmt.Sprintf("opened port %s on instance %s", port, instanceName))
}

func ClosePort() {
	co := resource.NewCloudOperator()

	port := args.Poll()
	instanceName := args.Poll()

	_ = ValidatePort(port)

	instance := co.Finder.FindInstanceByName(instanceName)

	resource.RemoveSecurityGroup(instance, port)

	fmt.Println(fmt.Sprintf("closed port %s on instance %s", port, instanceName))
}
