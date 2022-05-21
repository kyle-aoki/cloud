package defaults

import (
	"fmt"

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

type CloudLabDefaultsOperator struct {
	Vpc             *ec2.Vpc
	PublicSubnet    *ec2.Subnet
	PrivateSubnet   *ec2.Subnet
	RouteTable      *ec2.RouteTable
	InternetGateway *ec2.InternetGateway
	SecurityGroups  []*ec2.SecurityGroup
	KeyPairs        []*ec2.KeyPairInfo
	CurrentKeyPair  *ec2.KeyPairInfo
	Instances       []*ec2.Instance
}

func NewOperator() *CloudLabDefaultsOperator {
	return &CloudLabDefaultsOperator{}
}

func Start() *CloudLabDefaultsOperator {
	cldo := NewOperator()
	cldo.FindAll()
	return cldo
}

func (cldo *CloudLabDefaultsOperator) FindAll() {
	cldo.findCloudLabVpc()
	cldo.findSubnets()
	cldo.findCloudLabRouteTable()
	cldo.findInternetGateway()
	cldo.findCloudLabSecurityGroups()
	cldo.FindAllInstances()
}

// Idempotent
func CreateCloudLabDefaults() {
	cldo := &CloudLabDefaultsOperator{}
	cldo.FindAll()

	cldo.createVpc()
	cldo.findCloudLabRouteTable()

	cldo.createSubnets()

	cldo.createInternetGateway()
	cldo.attachInternetGateway()

	cldo.addInternetGatewayRoute()
	cldo.associatePublicSubnetWithRouteTable()

	cldo.CreateSecurityGroup("22", 22)

	fmt.Println("all cloudlab resources exist")
	fmt.Println("to create a key-pair, run: 'cloudlab create key-pair'")
}

func DestroyCloudLabResources() {
	cldo := &CloudLabDefaultsOperator{}
	cldo.FindAll()

	if cldo.PublicIpAddressesExist() {
		panic("run 'lab delete all instances' and try again")
	}

	cldo.detachInternetGateway()
	cldo.deleteInternetGateway()
	cldo.deleteSubnets()
	cldo.deleteAllSecurityGroups()
	cldo.deleteVpc()

	fmt.Println("no cloudlab resources exist")
}
