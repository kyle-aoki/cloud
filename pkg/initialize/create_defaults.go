package initialize

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
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

func CreateCloudLabDefaults() {
	exists, vpcId := CloudLabVpcExists()
	if !exists {
		vpcId = CreateCloudLabVpc()
	}
	exists, subnetId := PublicSubnetExists(vpcId)
	if !exists {
		subnetId = CreatePublicSubnet(vpcId)
		ModifyPublicSubnetAttributes(subnetId)
	}
	if !CloudLabRouteTableExists() {
		CreateCloudLabRouteTable(vpcId)
	}
}

func CreateCloudLabVpc() *string {
	fmt.Println("creating cloudlab VPC")
	cvo, err := amazon.EC2().CreateVpc(&ec2.CreateVpcInput{
		CidrBlock:         util.StrPtr(DefaultVpcCidrBlock),
		TagSpecifications: NameTag("vpc", DefaultVpcName),
	})
	util.MustExec(err)

	return cvo.Vpc.VpcId
}

func CreatePublicSubnet(vpcId *string) *string {
	fmt.Println("creating public subnet")
	cso, err := amazon.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		VpcId:             vpcId,
		CidrBlock:         util.StrPtr(PublicSubnetCidrBlock),
		TagSpecifications: NameTag("subnet", "cloudlab-public-subnet"),
	})
	util.MustExec(err)

	return cso.Subnet.SubnetId
}

func ModifyPublicSubnetAttributes(subnetId *string) {
	fmt.Println("modifying public subnet attributes")
	_, err := amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
		SubnetId:                             subnetId,
		EnableResourceNameDnsARecordOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
	})
	util.MustExec(err)

	_, err = amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
		SubnetId:            subnetId,
		MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
	})
	util.MustExec(err)
}

func PublicSubnetExists(vpcId *string) (exists bool, subnetId *string) {
	err := amazon.EC2().DescribeSubnetsPages(
		&ec2.DescribeSubnetsInput{},
		func(dso *ec2.DescribeSubnetsOutput, b bool) bool {
			for _, subnet := range dso.Subnets {
				if *subnet.VpcId == *vpcId {
					util.VPrint("found cloudlab public subnet", *subnet.SubnetId)
					exists = true
					subnetId = subnet.SubnetId
					return false
				}
			}
			return true
		},
	)
	if !exists {
		fmt.Println("did not find public subnet")
	}
	util.MustExec(err)
	return exists, subnetId
}

func CloudLabRouteTableExists() bool {
	var NextToken *string
	for {
		drto, err := amazon.EC2().DescribeRouteTables(&ec2.DescribeRouteTablesInput{
			NextToken: NextToken,
		})
		util.MustExec(err)
		for _, rt := range drto.RouteTables {
			for _, tag := range rt.Tags {
				if *tag.Key == "Name" && *tag.Value == CloudLabRouteTable {
					util.VPrint("found cloudlab route table", *rt.RouteTableId)
					return true
				}
			}
		}
		if drto.NextToken == nil {
			fmt.Println("did not find cloudlab route table")
			return false
		}
		NextToken = drto.NextToken
	}
}

func CreateCloudLabRouteTable(vpcId *string) {
	crto, err := amazon.EC2().CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId:             vpcId,
		TagSpecifications: NameTag("route-table", CloudLabRouteTable),
	})
	util.MustExec(err)
	util.VPrint("created cloudlab route table", *crto.RouteTable.RouteTableId)
}

func NameTag(resourceType string, name string) []*ec2.TagSpecification {
	return []*ec2.TagSpecification{{
		ResourceType: util.StrPtr(resourceType),
		Tags: []*ec2.Tag{{
			Key:   util.StrPtr("Name"),
			Value: util.StrPtr(name),
		}}},
	}
}
