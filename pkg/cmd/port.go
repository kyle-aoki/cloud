package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
)

// ####################################################################################
// ####################################################################################
// ####################################################################################

func OpenPorts() {
	instanceName := args.Poll()
	ports := args.Collect()

	for _, port := range ports {
		openPort(instanceName, port)
	}
}

func openPort(name string, port string) {
	util.Log("port %s", port)
	util.Log("instanceName %s", name)

	co := resource.New()
	co.Rs.Vpc = co.Finder.FindCloudLabVpcOrPanic()
	co.Rs.SecurityGroups = co.Finder.FindAllSecurityGroups()
	instance := co.Finder.FindInstanceByNameOrPanic(name)

	if !resource.SecurityGroupExists(co.Rs.SecurityGroups, port) {
		resource.CreateSecurityGroup(co.Rs.Vpc, port)
		co.Rs.SecurityGroups = co.Finder.FindAllSecurityGroups()
	}

	if resource.InstanceHasPortOpen(instance, port) {
		panic(fmt.Sprintf("port %s already open on instance %s\n", port, name))
	}

	securityGroup := resource.SecurityGroupByNameOrPanic(co.Rs.SecurityGroups, port)

	resource.OpenPort(instance, securityGroup)
	fmt.Println(fmt.Sprintf("opened port %s on instance %s", port, name))
}

// ####################################################################################
// ####################################################################################
// ####################################################################################

func ClosePorts() {
	instanceName := args.Poll()
	ports := args.Collect()
	for _, port := range ports {
		closePort(instanceName, port)
	}
}

func closePort(instanceName string, port string) {
	co := resource.New()
	co.Rs.Instances = co.Finder.FindNonTerminatedInstances()
	_ = resource.ValidatePort(port)
	instance := co.Finder.FindInstanceByName(instanceName)
	resource.ClosePort(instance, port)
	fmt.Println(fmt.Sprintf("closed port %s on instance %s", port, instanceName))
}

// ####################################################################################
// ####################################################################################
// ####################################################################################
