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

func (cldo *CloudLabDefaultsOperator) FindInstanceByName(name string) *ec2.Instance {
	for _, inst := range cldo.Instances {
		if NameTagEquals(inst.Tags, name) {
			return inst
		}
	}
	panic("node not found")
}

func (cldo *CloudLabDefaultsOperator) PublicIpAddressesExist() bool {
	for _, inst := range cldo.Instances {
		if inst.PublicIpAddress != nil {
			return true
		}
	}
	return false
}

func (cldo *CloudLabDefaultsOperator) GetInstanceByName(name string) *ec2.Instance {
	for _, inst := range cldo.Instances {
		if NameTagEquals(inst.Tags, name) {
			return inst
		}
	}
	panic("node not found")
}

func (cldo *CloudLabDefaultsOperator) AssignSecurityGroup(
	instance *ec2.Instance,
	securityGroup *ec2.SecurityGroup,
) {
	var groupIds []*string
	for _, sgs := range instance.SecurityGroups {
		groupIds = append(groupIds, sgs.GroupId)
	}
	groupIds = append(groupIds, securityGroup.GroupId)
	_, err := amazon.EC2().ModifyInstanceAttribute(&ec2.ModifyInstanceAttributeInput{
		InstanceId: instance.InstanceId,
		Groups:     groupIds,
	})
	util.MustExec(err)
}
