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
}

func (cldo *CloudLabDefaultsOperator) FindAll() {
	cldo.findCloudLabVpc()
	cldo.PublicSubnet = findSubnet(CloudLabPublicSubnetName)
	cldo.PrivateSubnet = findSubnet(CloudLabPrivateSubnetName)
	cldo.findCloudLabRouteTable()
	cldo.findInternetGateway()
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

	fmt.Println("all cloudlab resources exist")
}

func DestroyCloudLabResources() {
	cldo := &CloudLabDefaultsOperator{}
	cldo.FindAll()

	cldo.detachInternetGateway()
	cldo.deleteInternetGateway()
	cldo.deleteSubnets()
	cldo.deleteVpc()

	fmt.Println("no cloudlab resources exist")
}
