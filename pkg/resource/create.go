package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type ResourceCreator interface {
	createVpc(cidrBlock string, name string) *ec2.Vpc
	createSubnet(vpc *ec2.Vpc, name string, cidr string) *ec2.Subnet
	createInternetGateway(name string) *ec2.InternetGateway
	createRouteTable(vpc *ec2.Vpc, name string) *ec2.RouteTable
}

type AWSCreator struct{}

func (c *AWSCreator) createVpc(cidrBlock string, name string) *ec2.Vpc {
	cvo, err := amazon.EC2().CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: util.StrPtr(cidrBlock),

		TagSpecifications: CreateTagSpecs("vpc", map[string]string{
			"Name": name,
		}),
	})
	util.MustExec(err)
	return cvo.Vpc
}

func (c *AWSCreator) createSubnet(vpc *ec2.Vpc, name string, cidr string) *ec2.Subnet {
	cso, err := amazon.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		VpcId:             vpc.VpcId,
		CidrBlock:         util.StrPtr(cidr),
		TagSpecifications: CreateNameTagSpec("subnet", name),
	})
	util.MustExec(err)
	return cso.Subnet
}

func (c *AWSCreator) createInternetGateway(name string) *ec2.InternetGateway {
	cigo, err := amazon.EC2().CreateInternetGateway(&ec2.CreateInternetGatewayInput{
		TagSpecifications: CreateNameTagSpec("internet-gateway", name),
	})
	util.MustExec(err)
	return cigo.InternetGateway
}

func (c *AWSCreator) createRouteTable(vpc *ec2.Vpc, name string) *ec2.RouteTable {
	crto, err := amazon.EC2().CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId: vpc.VpcId,
		TagSpecifications: CreateTagSpecs("route-table", map[string]string{
			"Name": name,
		}),
	})
	util.MustExec(err)
	return crto.RouteTable
}
