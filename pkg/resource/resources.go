package resource

import "github.com/aws/aws-sdk-go/service/ec2"

type AWSResources struct {
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
