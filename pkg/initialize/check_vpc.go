package initialize

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func CloudLabVpcExists() (bool, *string) {
	dvo, err := amazon.EC2().DescribeVpcs(&ec2.DescribeVpcsInput{})
	util.MustExec(err)

	vpcNameTags := GetVpcNameTags(dvo)

	if util.Contains(vpcNameTags, DefaultVpcName) {
		vpcId := GetVpcIdByNameTag(dvo, DefaultVpcName)
		util.VPrint("found cloudlab vpc", *vpcId)
		return true, vpcId
	}

	fmt.Println("did not find cloudlab vpc")
	return false, nil
}

func GetVpcIdByNameTag(dvo *ec2.DescribeVpcsOutput, name string) *string {
	for _, vpc := range dvo.Vpcs {
		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				if *tag.Value == name {
					return vpc.VpcId
				}
			}
		}
	}
	panic("failed to find vpc id")
}

func GetVpcNameTags(dvo *ec2.DescribeVpcsOutput) []string {
	if dvo.NextToken != nil {
		log.Fatal("You have too many VPCs for this program to handle.")
	}
	var vpcNameTags []string
	for _, vpc := range dvo.Vpcs {
		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				vpcNameTags = append(vpcNameTags, *tag.Value)
			}
		}
	}
	return vpcNameTags
}

// (*ec2.DescribeVpcsOutput)(0xc000013360)({
// 	Vpcs: [{
// 		CidrBlock: "10.0.0.0/16",
// 		CidrBlockAssociationSet: [{
// 			AssociationId: "vpc-cidr-assoc-084e1cd5c75e33c4f",
// 			CidrBlock: "10.0.0.0/16",
// 			CidrBlockState: {
// 			  State: "associated"
// 			}
// 		  }],
// 		DhcpOptionsId: "dopt-04b5d348dcaa50388",
// 		InstanceTenancy: "default",
// 		IsDefault: false,
// 		OwnerId: "893748410716",
// 		State: "available",
// 		Tags: [{
// 			Key: "Name",
// 			Value: "main"
// 		  }],
// 		VpcId: "vpc-0a7dfa2f42329419e"
// 	  },{
// 		CidrBlock: "172.30.0.0/16",
// 		CidrBlockAssociationSet: [{
// 			AssociationId: "vpc-cidr-assoc-0717d595450e1be1b",
// 			CidrBlock: "172.30.0.0/16",
// 			CidrBlockState: {
// 			  State: "associated"
// 			}
// 		  }],
// 		DhcpOptionsId: "dopt-04b5d348dcaa50388",
// 		InstanceTenancy: "default",
// 		IsDefault: false,
// 		OwnerId: "893748410716",
// 		State: "available",
// 		VpcId: "vpc-0e38774c9ba92421b"
// 	  }]
//   })
