package cmd

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func DeleteInstances() {
	var targets []string = args.Collect()

	ro := resource.NewResourceOperator()
	var names []string
	var ids []*string

	for _, inst := range ro.Instances {
		name := resource.FindNameTagValue(inst.Tags)
		for _, target := range targets {
			if name != nil && *name == target {
				names = append(names, *name)
				ids = append(ids, inst.InstanceId)
			}
		}
	}

	_, err := amazon.EC2().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: ids,
	})
	util.MustExec(err)

	for _, name := range names {
		fmt.Println(fmt.Sprintf("deleted %s", name))
	}
}
