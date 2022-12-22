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
	instanceName := args.PollOrPanic()
	ports := args.CollectOrPanic()

	for _, port := range ports {
		openPort(instanceName, port)
	}
}

func openPort(name string, port string) {
	util.Log("port %s", port)
	util.Log("instanceName %s", name)

	lr := resource.NewLabResources()
	lr.Vpc = resource.FindCloudLabVpcOrPanic()
	lr.SecurityGroups = resource.FindAllSecurityGroups()
	instance := resource.FindInstanceByNameOrPanic(name)

	if !resource.SecurityGroupExists(lr.SecurityGroups, port) {
		resource.CreateSecurityGroup(lr.Vpc, port)
		lr.SecurityGroups = resource.FindAllSecurityGroups()
	}

	if resource.HasPortOpen(instance, port) {
		panic(fmt.Sprintf("port %s already open on instance %s\n", port, name))
	}

	securityGroup := resource.SecurityGroupByNameOrPanic(lr.SecurityGroups, port)

	resource.OpenPort(instance, securityGroup)
	fmt.Printf("opened port %s on instance %s\n", port, name)
}

// ####################################################################################
// ####################################################################################
// ####################################################################################

func ClosePorts() {
	instanceName := args.PollOrPanic()
	ports := args.CollectOrPanic()
	for _, port := range ports {
		closePort(instanceName, port)
	}
}

func closePort(instanceName string, port string) {
	lr := resource.NewLabResources()
	lr.Instances = resource.FindNonTerminatedInstances()
	_ = resource.ValidatePort(port)
	instance := resource.FindInstanceByName(instanceName)
	resource.ClosePort(instance, port)
	fmt.Printf("closed port %s on instance %s\n", port, instanceName)
}

// ####################################################################################
// ####################################################################################
// ####################################################################################
