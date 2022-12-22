package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func createVpc(cidrBlock string, name string) *ec2.Vpc {
	log.Println("creating vpc")
	cvo, err := amazon.EC2().CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: util.StrPtr(cidrBlock),
		TagSpecifications: CreateTagSpecs("vpc", map[string]string{
			"Name": name,
		}),
	})
	util.Check(err)
	return cvo.Vpc
}

func createSubnet(vpc *ec2.Vpc, name string, cidr string) *ec2.Subnet {
	log.Println("creating subnet")
	cso, err := amazon.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		VpcId:             vpc.VpcId,
		CidrBlock:         &cidr,
		TagSpecifications: CreateNameTagSpec("subnet", name),
	})
	util.Check(err)
	return cso.Subnet
}

func createInternetGateway(name string) *ec2.InternetGateway {
	log.Println("creating internet gateway")
	cigo, err := amazon.EC2().CreateInternetGateway(&ec2.CreateInternetGatewayInput{
		TagSpecifications: CreateNameTagSpec("internet-gateway", name),
	})
	util.Check(err)
	return cigo.InternetGateway
}

func createRouteTable(vpc *ec2.Vpc, name string) *ec2.RouteTable {
	log.Println("creating route table")
	crto, err := amazon.EC2().CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId: vpc.VpcId,
		TagSpecifications: CreateTagSpecs("route-table", map[string]string{
			"Name": name,
		}),
	})
	util.Check(err)
	return crto.RouteTable
}

func createKeyPair(name string) *string {
	log.Println("making create key pair request")
	ckpo, err := amazon.EC2().CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName:           &name,
		TagSpecifications: CreateNameTagSpec("key-pair", name),
	})
	util.Check(err)
	return ckpo.KeyMaterial
}

func createInboundRule(groupId *string, protocol Protocol, port int) {
	_, err := amazon.EC2().AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:    groupId,
		FromPort:   util.Ptr(int64(port)),
		ToPort:     util.Ptr(int64(port)),
		IpProtocol: util.Ptr(string(protocol)),
		CidrIp:     util.Ptr(AllIpsCidr),
	})
	util.Check(err)
}

func CreateSecurityGroup(vpc *ec2.Vpc, port string) {
	portInt := ValidatePort(port)
	csgo, err := amazon.EC2().CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		VpcId:             vpc.VpcId,
		GroupName:         &port,
		Description:       &port,
		TagSpecifications: CreateNameTagSpec("security-group", CloudLabSecutiyGroup),
	})
	util.Check(err)
	createInboundRule(csgo.GroupId, "tcp", portInt)
	createInboundRule(csgo.GroupId, "udp", portInt)
}
