package resource

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Create cloudlab VPC if not exists
// Create Public Subnet if not exists
// Modify Public Subnet Attributes

// Create Route Table ON a VPC
// Create Internet Gateway
// Attach Internet Gateway to VPC
// Set up 0.0.0.0/0 --> IGW on Route Table
// Set up Route Table Subnet Association

type ResourceOperator struct {
	Vpc               *ec2.Vpc
	PublicSubnet      *ec2.Subnet
	PrivateSubnet     *ec2.Subnet
	PublicRouteTable  *ec2.RouteTable
	PrivateRouteTable *ec2.RouteTable
	InternetGateway   *ec2.InternetGateway
	SecurityGroups    []*ec2.SecurityGroup
	KeyPairs          []*ec2.KeyPairInfo
	Instances         []*ec2.Instance
	CurrentKeyPair    *ec2.KeyPairInfo
}

func New() *ResourceOperator {
	return &ResourceOperator{}
}

func NewResourceOperator() *ResourceOperator {
	ro := New()
	ro.FindAll()
	ro.Audit()
	return ro
}

func FatalMissing(resourceType string) {
	fmt.Printf("%s is missing from cloudlab resources\n", resourceType)
	fmt.Println("run 'lab init' to create missing cloudlab resources")
	os.Exit(1)
}

func (ro *ResourceOperator) FindAll() {
	ro.Vpc = findVpc(CloudLabVpc)
	if ro.Vpc == nil {
		return
	}
	ro.PublicSubnet = findSubnet(CloudLabPublicSubnet)
	ro.PrivateSubnet = findSubnet(CloudLabPrivateSubnet)
	ro.PublicRouteTable = findMainRouteTable(ro.Vpc)
	ro.PrivateRouteTable = findRouteTable(*ro.Vpc.VpcId, CloudLabPrivateRouteTable)
	ro.InternetGateway = findInternetGateway()
	ro.SecurityGroups = findCloudLabSecurityGroups()
	ro.Instances = findInstances()
	ro.KeyPairs = FindAllCloudLabKeyPairs()
	ro.CurrentKeyPair = FindCurrentKeyPair(ro.KeyPairs)
}

func (ro *ResourceOperator) Audit() {
	if ro.Vpc == nil {
		FatalMissing("vpc")
	}
	if ro.PublicSubnet == nil {
		FatalMissing("public subnet")
	}
	if ro.PrivateSubnet == nil {
		FatalMissing("private subnet")
	}
	if ro.PublicRouteTable == nil {
		FatalMissing("main route table")
	}
	if ro.PrivateRouteTable == nil {
		FatalMissing("private route table")
	}
	if ro.InternetGateway == nil {
		FatalMissing("internet gateway")
	}
	if ro.SecurityGroups == nil {
		FatalMissing("security groups")
	}
	if ro.KeyPairs == nil {
		FatalMissing("cloudlab key pair")
	}
	if ro.CurrentKeyPair == nil {
		FatalMissing("cloudlab key pair")
	}
}

func (ro *ResourceOperator) InitializeCloudLabResources() {
	if ro.Vpc == nil {
		ro.Vpc = createVpc(DefaultVpcCidrBlock, CloudLabVpc)
		ro.PublicRouteTable = findMainRouteTable(ro.Vpc)
	}

	if ro.PublicSubnet == nil {
		ro.PublicSubnet = createSubnet(ro.Vpc, CloudLabPublicSubnet, PublicSubnetCidrBlock)
	}
	if ro.PrivateSubnet == nil {
		ro.PrivateSubnet = createSubnet(ro.Vpc, CloudLabPrivateSubnet, PrivateSubnetCidrBlock)
	}

	if ro.PrivateRouteTable == nil {
		ro.PrivateRouteTable = createRouteTable(ro.Vpc, CloudLabPrivateRouteTable)
	}

	if ro.InternetGateway == nil {
		ro.InternetGateway = createInternetGateway(CloudLabInternetGateway)
	}

	if !internetGatewayIsAttachedToVpc(ro.InternetGateway, ro.Vpc) {
		attachInternetGateway(ro.InternetGateway, ro.Vpc)
	}

	if !internetGatewayRouteExistsOnRouteTable(ro.PublicRouteTable, ro.InternetGateway) {
		addInternetGatewayRoute(ro.PublicRouteTable, ro.InternetGateway, RouteTablePublicSubnetCidr)
	}

	if !subnetAssociationExistsOnRouteTable(ro.PublicRouteTable, ro.PublicSubnet) {
		associateSubnetWithRouteTable(ro.PublicRouteTable, ro.PublicSubnet)
	}

	securityGroup22 := findSecurityGroupByName(ro.SecurityGroups, "22")
	if securityGroup22 == nil {
		CreateSecurityGroup(ro.Vpc, "22", 22)
		ro.SecurityGroups = findCloudLabSecurityGroups()
	}
}

func (ro *ResourceOperator) DestroyCloudLabResources() {
	if len(findNotTerminatedInstances(ro.Instances)) > 0 {
		panic("run 'lab delete all instances' and try again")
	}

	if len(ro.Instances) > 0 {
		deleteInstances(ro.Instances)
	}
	if ro.InternetGateway != nil {
		detachInternetGateway(ro.InternetGateway, ro.Vpc)
		deleteInternetGateway(ro.InternetGateway)
	}

	if ro.PrivateRouteTable != nil {
		disassociateSubnetsFromRouteTable(ro.PrivateRouteTable)
		deleteRouteTable(ro.PrivateRouteTable)
	}

	if ro.PublicSubnet != nil {
		deleteSubnet(ro.PublicSubnet)
	}
	if ro.PrivateSubnet != nil {
		deleteSubnet(ro.PrivateSubnet)
	}

	deleteSecurityGroups(ro.SecurityGroups)

	if ro.Vpc != nil {
		deleteVpc(ro.Vpc)
	}

	deleteKeyPairs(ro.KeyPairs)
}
