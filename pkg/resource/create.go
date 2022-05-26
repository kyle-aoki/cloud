package resource

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func createVpc(cidrBlock string, name string) *ec2.Vpc {
	cvo, err := amazon.EC2().CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: util.StrPtr(cidrBlock),
		
		TagSpecifications: CreateTagSpecs("vpc", map[string]string{
			"Name": name,
		}),
	})
	util.MustExec(err)
	return cvo.Vpc
}

func createSubnet(vpcId string, name string, cidr string) (sn *ec2.Subnet) {
	cso, err := amazon.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		VpcId:             util.StrPtr(vpcId),
		CidrBlock:         util.StrPtr(cidr),
		TagSpecifications: CreateNameTagSpec("subnet", name),
	})
	util.MustExec(err)
	return cso.Subnet
}

func createInternetGateway(name string) *ec2.InternetGateway {
	cigo, err := amazon.EC2().CreateInternetGateway(&ec2.CreateInternetGatewayInput{
		TagSpecifications: CreateNameTagSpec("internet-gateway", name),
	})
	util.MustExec(err)
	return cigo.InternetGateway
}

func createRouteTable(vpc *ec2.Vpc, name string) *ec2.RouteTable {
	crto, err := amazon.EC2().CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId: vpc.VpcId,
		TagSpecifications: CreateTagSpecs("route-table", map[string]string{
			"Name": name,
		}),
	})
	util.MustExec(err)
	return crto.RouteTable
}
