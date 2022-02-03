package delete

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/get"
	"cloud/pkg/tags"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func DeleteNodes() {
	names := args.Collect()
	nodes := get.GetNodes()
	var instanceIds []*string

	for _, node := range nodes {
		nodeName := tags.GetName(node)
		for _, name := range names {
			if nodeName == name {
				instanceIds = append(instanceIds, node.InstanceId)
			}
		}
	}

	tio, err := amazon.EC2Client().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	})
	util.Check(err)

	for _, ti := range tio.TerminatingInstances {
		fmt.Println(*ti.InstanceId)
	}
}
