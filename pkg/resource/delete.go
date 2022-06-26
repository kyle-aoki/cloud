package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type ResourceDeleter struct{}

func (a *ResourceDeleter) deleteSubnet(subnet *ec2.Subnet) {
	log.Println("deleting subnet")
	_, err := amazon.EC2().DeleteSubnet(&ec2.DeleteSubnetInput{
		SubnetId: subnet.SubnetId,
	})
	util.MustExec(err)
}

func (a *ResourceDeleter) deleteInternetGateway(ig *ec2.InternetGateway) {
	log.Println("deleting internet gateway")
	_, err := amazon.EC2().DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
		InternetGatewayId: ig.InternetGatewayId,
	})
	util.MustExec(err)
}

func (a *ResourceDeleter) deleteInstance(instance *ec2.Instance) {
	log.Println("deleting instance")
	_, err := amazon.EC2().TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{instance.InstanceId},
	})
	util.MustExec(err)
}

func (a *ResourceDeleter) deleteInstances(instances []*ec2.Instance) {
	for _, instance := range instances {
		a.deleteInstance(instance)
	}
}

func (a *ResourceDeleter) deleteSecurityGroup(sg *ec2.SecurityGroup) {
	log.Println("deleting security group")
	if sg.GroupName != nil && *sg.GroupName == "Default" {
		return
	}
	_, err := amazon.EC2().DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
		GroupId: sg.GroupId,
	})
	util.MustExec(err)
}

func (a *ResourceDeleter) deleteSecurityGroups(sgs []*ec2.SecurityGroup) {
	for _, sg := range sgs {
		a.deleteSecurityGroup(sg)
	}
}

func (a *ResourceDeleter) deleteVpc(vpc *ec2.Vpc) {
	log.Println("deleting vpc")
	_, err := amazon.EC2().DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: vpc.VpcId,
	})
	util.MustExec(err)
}

func (a *ResourceDeleter) deleteKeyPair(key *ec2.KeyPairInfo) {
	log.Println("deleting key pair info")
	_, err := amazon.EC2().DeleteKeyPair(&ec2.DeleteKeyPairInput{
		KeyPairId: key.KeyPairId,
	})
	util.MustExec(err)
}

func (a *ResourceDeleter) deleteKeyPairs(keys []*ec2.KeyPairInfo) {
	for _, key := range keys {
		a.deleteKeyPair(key)
	}
}

func (a *ResourceDeleter) deleteRouteTable(rt *ec2.RouteTable) {
	log.Println("deleting route table")
	_, err := amazon.EC2().DeleteRouteTable(&ec2.DeleteRouteTableInput{
		RouteTableId: rt.RouteTableId,
	})
	util.MustExec(err)
}
