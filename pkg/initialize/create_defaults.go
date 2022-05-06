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
	if isMissingCloudLabVpc() {
		CreateCloudLabVpc()
	}
	vpcId := findCloudLabVpcId()

	if isMissingCloudLabPublicSubnet(vpcId) {
		CreatePublicSubnet(vpcId)
	}
	subnetId := findSubnetByVpcId(vpcId)

	ModifyPublicSubnetAttributes(subnetId)

	if !CloudLabRouteTableExists() {
		CreateCloudLabRouteTable(vpcId)
	}
}

func isMissingCloudLabVpc() bool {
	vpcId := findCloudLabVpcId()
	if vpcId == nil {
		fmt.Println("missing cloudlab vpc")
		return true
	}
	util.VPrint("found cloudlab vpc", *vpcId)
	return false
}

func findCloudLabVpcId() *string {
	var NextToken *string
	for {
		dvo, err := amazon.EC2().DescribeVpcs(&ec2.DescribeVpcsInput{NextToken: NextToken})
		util.MustExec(err)
		for _, vpc := range dvo.Vpcs {
			nameTagValue := findNameTagValue(vpc.Tags)
			if nameTagValue != nil && *nameTagValue == DefaultVpcName {
				return vpc.VpcId
			}
		}
		NextToken = dvo.NextToken
		if NextToken == nil {
			return nil
		}
	}
}

func isMissingPublicSubnetAttributes() bool {
	var NextToken *string
	for {
		dso, err := amazon.EC2().DescribeSubnets(&ec2.DescribeSubnetsInput{NextToken: NextToken})
		util.MustExec(err)
		for _, subnet := range dso.Subnets {
			nameTagValue := findNameTagValue(subnet.Tags)
			if nameTagValue != nil && *nameTagValue == CloudLabPublicSubnetNameTagValue {
				if subnet.MapPublicIpOnLaunch == nil || !*subnet.MapPublicIpOnLaunch {
					return true
				}
				return false
			}
		}
		NextToken = dso.NextToken
		if NextToken == nil {
			return true
		}
	}
}
// MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},

func CreateCloudLabVpc() {
	fmt.Println("creating cloudlab VPC")
	_, err := amazon.EC2().CreateVpc(&ec2.CreateVpcInput{
		CidrBlock:         util.StrPtr(DefaultVpcCidrBlock),
		TagSpecifications: NameTag("vpc", DefaultVpcName),
	})
	util.MustExec(err)
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
	// _, err := amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
	// 	SubnetId:                             subnetId,
	// 	EnableResourceNameDnsARecordOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
	// })
	// util.MustExec(err)

	// _, err = amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
	// 	SubnetId:            subnetId,
	// 	MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
	// })
	// util.MustExec(err)
}

func findSubnetByVpcId(vpcId *string) *string {
	var NextToken *string
	for {
		dso, err := amazon.EC2().DescribeSubnets(&ec2.DescribeSubnetsInput{NextToken: NextToken})
		util.MustExec(err)
		for _, subnet := range dso.Subnets {
			if subnet.VpcId != nil && *subnet.VpcId == *vpcId {
				return subnet.VpcId
			}
		}
		NextToken = dso.NextToken
		if NextToken == nil {
			return nil
		}
	}
}

func isMissingCloudLabPublicSubnet(vpcId *string) bool {
	subnetId := findSubnetByVpcId(vpcId)
	if subnetId == nil {
		fmt.Println("missing cloudlab public subnet")
		return true
	}
	util.VPrint("found cloudlab public subnet", *subnetId)
	return false
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
