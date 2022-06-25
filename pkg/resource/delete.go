package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type ResourceDeleter interface {
	deleteSubnet(subnet *ec2.Subnet)
	deleteInternetGateway(ig *ec2.InternetGateway)
	deleteInstance(instance *ec2.Instance)
	deleteInstances(instances []*ec2.Instance)
	deleteSecurityGroup(sg *ec2.SecurityGroup)
	deleteSecurityGroups(sgs []*ec2.SecurityGroup)
	deleteVpc(vpc *ec2.Vpc)
	deleteKeyPair(key *ec2.KeyPairInfo)
	deleteKeyPairs(keys []*ec2.KeyPairInfo)
	deleteRouteTable(rt *ec2.RouteTable)
}

type AWSDeleter struct{}

func (a *AWSDeleter) deleteSubnet(subnet *ec2.Subnet) {
	_, err := amazon.EC2().DeleteSubnet(&ec2.DeleteSubnetInput{
		SubnetId: subnet.SubnetId,
	})
	util.MustExec(err)
}

func (a *AWSDeleter) deleteInternetGateway(ig *ec2.InternetGateway) {
	_, err := amazon.EC2().DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
		InternetGatewayId: ig.InternetGatewayId,
	})
	util.MustExec(err)
}

func (a *AWSDeleter) deleteInstance(instance *ec2.Instance) {
	_, err := amazon.EC2().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{instance.InstanceId},
	})
	util.MustExec(err)
}

func (a *AWSDeleter) deleteInstances(instances []*ec2.Instance) {
	for _, instance := range instances {
		a.deleteInstance(instance)
	}
}

func (a *AWSDeleter) deleteSecurityGroup(sg *ec2.SecurityGroup) {
	if sg.GroupName != nil && *sg.GroupName == "Default" {
		return
	}
	_, err := amazon.EC2().DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
		GroupId: sg.GroupId,
	})
	util.MustExec(err)
}

func (a *AWSDeleter) deleteSecurityGroups(sgs []*ec2.SecurityGroup) {
	for _, sg := range sgs {
		a.deleteSecurityGroup(sg)
	}
}

func (a *AWSDeleter) deleteVpc(vpc *ec2.Vpc) {
	_, err := amazon.EC2().DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: vpc.VpcId,
	})
	util.MustExec(err)
}

func (a *AWSDeleter) deleteKeyPair(key *ec2.KeyPairInfo) {
	_, err := amazon.EC2().DeleteKeyPair(&ec2.DeleteKeyPairInput{
		KeyPairId: key.KeyPairId,
	})
	util.MustExec(err)
}

func (a *AWSDeleter) deleteKeyPairs(keys []*ec2.KeyPairInfo) {
	for _, key := range keys {
		a.deleteKeyPair(key)
	}
}

func (a *AWSDeleter) deleteRouteTable(rt *ec2.RouteTable) {
	_, err := amazon.EC2().DeleteRouteTable(&ec2.DeleteRouteTableInput{
		RouteTableId: rt.RouteTableId,
	})
	util.MustExec(err)
}
