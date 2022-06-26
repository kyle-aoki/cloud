package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type ResourceCreator struct{}

func (c *ResourceCreator) createVpc(cidrBlock string, name string) *ec2.Vpc {
	cvo, err := amazon.EC2().CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: util.StrPtr(cidrBlock),

		TagSpecifications: CreateTagSpecs("vpc", map[string]string{
			"Name": name,
		}),
	})
	util.MustExec(err)
	return cvo.Vpc
}

func (c *ResourceCreator) createSubnet(vpc *ec2.Vpc, name string, cidr string) *ec2.Subnet {
	cso, err := amazon.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		VpcId:             vpc.VpcId,
		CidrBlock:         util.StrPtr(cidr),
		TagSpecifications: CreateNameTagSpec("subnet", name),
	})
	util.MustExec(err)
	return cso.Subnet
}

func (c *ResourceCreator) createInternetGateway(name string) *ec2.InternetGateway {
	cigo, err := amazon.EC2().CreateInternetGateway(&ec2.CreateInternetGatewayInput{
		TagSpecifications: CreateNameTagSpec("internet-gateway", name),
	})
	util.MustExec(err)
	return cigo.InternetGateway
}

func (c *ResourceCreator) createRouteTable(vpc *ec2.Vpc, name string) *ec2.RouteTable {
	crto, err := amazon.EC2().CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId: vpc.VpcId,
		TagSpecifications: CreateTagSpecs("route-table", map[string]string{
			"Name": name,
		}),
	})
	util.MustExec(err)
	return crto.RouteTable
}

func (c *ResourceCreator) ExecuteCreateKeyPairRequest(name string) *string {
	log.Println("making create key pair request")
	ckpo, err := amazon.EC2().CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName:           util.StrPtr(name),
		TagSpecifications: CreateNameTagSpec("key-pair", name),
	})
	util.MustExec(err)
	return ckpo.KeyMaterial
}
