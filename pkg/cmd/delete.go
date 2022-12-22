package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func DeleteInstances() {
	targets := args.CollectOrPanic()
	util.Log("found delete targets: %v", targets)

	lr := resource.NewLabResources()
	lr.PublicSubnet = resource.FindPublicSubnet()
	lr.PrivateSubnet = resource.FindPrivateSubnet()
	lr.Instances = resource.FindNonTerminatedInstances()

	var targetInstances []*ec2.Instance

	for i := 0; i < len(lr.Instances); i++ {
		instName := resource.FindNameTagValue(lr.Instances[i].Tags)
		util.Log("found instance: ", lr.Instances[i].InstanceId)
		if instName != nil && util.Contains(*instName, targets) {
			targetInstances = append(targetInstances, lr.Instances[i])
		}
	}

	if len(targetInstances) == 0 {
		os.Exit(0)
	}

	for i := 0; i < len(targetInstances); i++ {
		util.Log("attempting to delete instance: %v", targetInstances[i].InstanceId)
	}

	resource.TerminateInstances(targetInstances)

	for i := 0; i < len(targetInstances); i++ {
		var ip *string
		if resource.InPrivateSubnet(targetInstances[i], lr) {
			ip = targetInstances[i].PrivateIpAddress
		} else {
			ip = targetInstances[i].PublicIpAddress
		}
		RemoveInstanceFromSshConfig(resource.FindNameTagValue(targetInstances[i].Tags), ip)
	}

	for i := 0; i < len(targetInstances); i++ {
		v := resource.FindNameTagValue(targetInstances[i].Tags)
		if v != nil {
			fmt.Println(*v)
		}
	}
}
