package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/defaults"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func DeleteInstances() {
	namesOfInstancesToDelete := args.Collect()

	cldo := defaults.Start()
	var instanceIds []*string

	for _, inst := range cldo.Instances {
		instanceName := defaults.FindNameTagValue(inst.Tags)
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
		util.VMessage("deleted", defaults.CloudLabInstance, *ti.InstanceId)
	}
}
