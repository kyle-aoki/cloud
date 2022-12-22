package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func deleteSubnet(subnet *ec2.Subnet) {
	util.Log("deleting subnet %s", *subnet.SubnetId)
	_, err := amazon.EC2().DeleteSubnet(&ec2.DeleteSubnetInput{
		SubnetId: subnet.SubnetId,
	})
	util.Check(err)
}

func deleteInternetGateway(ig *ec2.InternetGateway) {
	util.Log("deleting internet gateway %s", *ig.InternetGatewayId)
	_, err := amazon.EC2().DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
		InternetGatewayId: ig.InternetGatewayId,
	})
	util.Check(err)
}

func deleteSecurityGroup(sg *ec2.SecurityGroup) {
	util.Log("deleting security group %s", *sg.GroupId)
	if sg.GroupName != nil && *sg.GroupName == "Default" {
		return
	}
	_, err := amazon.EC2().DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
		GroupId: sg.GroupId,
	})
	util.Check(err)
}

func deleteSecurityGroups(sgs []*ec2.SecurityGroup) {
	for _, sg := range sgs {
		deleteSecurityGroup(sg)
	}
}

func deleteVpc(vpc *ec2.Vpc) {
	util.Log("deleting vpc %s", *vpc.VpcId)
	_, err := amazon.EC2().DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: vpc.VpcId,
	})
	util.Check(err)
}

func deleteKeyPair(key *ec2.KeyPairInfo) {
	util.Log("deleting key pair info %s", *key.KeyPairId)
	_, err := amazon.EC2().DeleteKeyPair(&ec2.DeleteKeyPairInput{
		KeyPairId: key.KeyPairId,
	})
	util.Check(err)
}

func deleteRouteTable(rt *ec2.RouteTable) {
	util.Log("deleting route table %s", *rt.RouteTableId)
	_, err := amazon.EC2().DeleteRouteTable(&ec2.DeleteRouteTableInput{
		RouteTableId: rt.RouteTableId,
	})
	util.Check(err)
}
