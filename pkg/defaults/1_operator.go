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
}

func NewOperator() *CloudLabDefaultsOperator {
	return &CloudLabDefaultsOperator{}
}

func (cldo *CloudLabDefaultsOperator) FindAll() {
	cldo.findCloudLabVpc()
	cldo.PublicSubnet = findSubnet(CloudLabPublicSubnetName)
	cldo.PrivateSubnet = findSubnet(CloudLabPrivateSubnetName)
	cldo.findCloudLabRouteTable()
	cldo.findInternetGateway()
	cldo.findCloudLabSecurityGroups()
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

	cldo.createSecurityGroup("allow-port-22", 22)
	cldo.createSecurityGroup("allow-port-3000", 3000)
	cldo.createSecurityGroup("allow-port-8080", 8080)

	fmt.Println("all cloudlab resources exist")
	fmt.Println("to create a key-pair, run: 'cloudlab key'")
}

func DestroyCloudLabResources() {
	cldo := &CloudLabDefaultsOperator{}
	cldo.FindAll()

	cldo.detachInternetGateway()
	cldo.deleteInternetGateway()
	cldo.deleteSubnets()
	cldo.deleteAllSecurityGroups()
	cldo.deleteVpc()

	fmt.Println("no cloudlab resources exist")
}
