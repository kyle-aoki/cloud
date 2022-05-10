package defaults

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (cldo *CloudLabDefaultsOperator) FindAllInstances() {
	client := amazon.EC2()

	di, err := client.DescribeInstances(&ec2.DescribeInstancesInput{})
	util.MustExec(err)

	var nodes []*ec2.Instance
	for _, res := range di.Reservations {
		for _, inst := range res.Instances {
			for _, tag := range inst.Tags {
				if tag.Key != nil && *tag.Key == CloudLabInstance {
					nodes = append(nodes, inst)
				}
			}
		}
	}
	cldo.Instances = nodes
}
