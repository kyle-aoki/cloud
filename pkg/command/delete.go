package cmd

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func DeleteInstances() {
	namesOfInstancesToDelete := args.Collect()

	ro := resource.NewResourceOperator()
	var instanceIds []*string

	for _, inst := range ro.Instances {
		instanceName := resource.FindNameTagValue(inst.Tags)
		for _, nameOfInstanceToDelete := range namesOfInstancesToDelete {
			if instanceName != nil && *instanceName == nameOfInstanceToDelete {
				instanceIds = append(instanceIds, inst.InstanceId)
			}
		}
	}

	tio, err := amazon.EC2().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	})
	util.MustExec(err)

	for _, ti := range tio.TerminatingInstances {
		util.VMessage("deleted", resource.CloudLabInstance, *ti.InstanceId)
	}
}

func DeleteAllInstances() {
	ro := resource.NewResourceOperator()
	var instanceIds []*string

	for _, inst := range ro.Instances {
		instanceIds = append(instanceIds, inst.InstanceId)
	}

	tio, err := amazon.EC2().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	})
	util.MustExec(err)

	for _, ti := range tio.TerminatingInstances {
		util.VMessage("deleted", resource.CloudLabInstance, *ti.InstanceId)
	}
}
