package resource

import (
	"cloudlab/pkg/util"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// ######################################################################################
// ######################################################################################
// ######################################################################################

type LabResources struct {
	Vpc               *ec2.Vpc
	PublicSubnet      *ec2.Subnet
	PrivateSubnet     *ec2.Subnet
	PublicRouteTable  *ec2.RouteTable
	PrivateRouteTable *ec2.RouteTable
	InternetGateway   *ec2.InternetGateway
	SecurityGroups    []*ec2.SecurityGroup
	Instances         []*ec2.Instance
	KeyPair           *ec2.KeyPairInfo
}

func NewLabResources() *LabResources {
	return &LabResources{}
}

// #############################################################################
// #############################################################################
// #############################################################################

func FindAllLabResources() *LabResources {
	log.Println("finding cloudlab resources...")
	lr := &LabResources{}
	lr.Vpc = FindVpc()
	if lr.Vpc == nil {
		return lr
	}
	lr.PublicSubnet = FindPublicSubnet()
	lr.PrivateSubnet = FindPrivateSubnet()
	lr.PublicRouteTable = findMainRouteTable(lr.Vpc)
	lr.PrivateRouteTable = findRouteTable(lr.Vpc, CloudLabPrivateRouteTable)
	lr.InternetGateway = findInternetGateway(CloudLabInternetGateway)
	lr.SecurityGroups = FindAllSecurityGroups()
	lr.Instances = FindInstances()
	lr.KeyPair = findKeyPair()
	return lr
}

// #############################################################################
// #############################################################################
// #############################################################################

// 1.  create key pair
// 2.  create vpc (auto creates public route table)
// 3.  create public subnet
// 4.  modify public subnet to enable IPv4 assignment
// 5.  create private subnet
// 6.  create private route table
// 7.  create internet gateway
// 8.  attach internet gateway to vpc
// 9.  add internet gateway to public route table
// 10. associate public route table with public subnet
// 11. create base security group for port 22 (tcp & udp)

func CreateMissingResources(lr *LabResources) {
	if lr.KeyPair == nil {
		ExecuteKeyPairCreationProcess()
		lr.KeyPair = findKeyPair()
	}
	if lr.Vpc == nil {
		lr.Vpc = createVpc(DefaultVpcCidrBlock, CloudLabVpc)
		lr.PublicRouteTable = findMainRouteTable(lr.Vpc)
	}
	if lr.PublicSubnet == nil {
		lr.PublicSubnet = createSubnet(lr.Vpc, CloudLabPublicSubnet, PublicSubnetCidrBlock)
		resolvePublicSubnetAttributes(lr.PublicSubnet)
	}
	if lr.PrivateSubnet == nil {
		lr.PrivateSubnet = createSubnet(lr.Vpc, CloudLabPrivateSubnet, PrivateSubnetCidrBlock)
	}
	if lr.PrivateRouteTable == nil {
		lr.PrivateRouteTable = createRouteTable(lr.Vpc, CloudLabPrivateRouteTable)
	}
	if lr.InternetGateway == nil {
		lr.InternetGateway = createInternetGateway(CloudLabInternetGateway)
	}
	if !internetGatewayIsAttachedToVpc(lr.InternetGateway, lr.Vpc) {
		attachInternetGatewayToVpc(lr.InternetGateway, lr.Vpc)
	}
	if !internetGatewayRouteExistsOnRouteTable(lr.PublicRouteTable, lr.InternetGateway) {
		addInternetGatewayRoute(lr.PublicRouteTable, lr.InternetGateway, RouteTablePublicSubnetCidr)
	}
	if !subnetAssociationExistsOnRouteTable(lr.PublicRouteTable, lr.PublicSubnet) {
		associateSubnetWithRouteTable(lr.PublicRouteTable, lr.PublicSubnet)
	}
	securityGroup22 := findSecurityGroupByName(lr.SecurityGroups, "22")
	if securityGroup22 == nil {
		CreateSecurityGroup(lr.Vpc, "22")
		lr.SecurityGroups = FindAllSecurityGroups()
	}
}

// #############################################################################
// #############################################################################
// #############################################################################

func DestroyCloudLabResources(lr *LabResources) {
	nonTerminatedInstances := FindNonTerminatedInstances()
	if len(nonTerminatedInstances) > 0 {
		panic("delete all instances before destroying cloudlab infrastructure")
	}
	if lr.InternetGateway != nil {
		detachInternetGatewayFromVpc(lr.InternetGateway, lr.Vpc)
		deleteInternetGateway(lr.InternetGateway)
	}
	if lr.PrivateRouteTable != nil {
		disassociateSubnetsFromRouteTable(lr.PrivateRouteTable)
		deleteRouteTable(lr.PrivateRouteTable)
	}
	if lr.PublicSubnet != nil {
		deleteSubnet(lr.PublicSubnet)
	}
	if lr.PrivateSubnet != nil {
		deleteSubnet(lr.PrivateSubnet)
	}
	deleteSecurityGroups(lr.SecurityGroups)
	if lr.Vpc != nil {
		deleteVpc(lr.Vpc)
	}
	if lr.KeyPair != nil {
		deleteKeyPair(lr.KeyPair)
	}
	if util.ObjectExists(KeyFilePath()) {
		err := os.Remove(KeyFilePath())
		if err != nil {
			panic("failed to delete key file: " + KeyFilePath())
		}
	}
}
