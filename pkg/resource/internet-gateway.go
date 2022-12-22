package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

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

func internetGatewayIsAttachedToVpc(ig *ec2.InternetGateway, vpc *ec2.Vpc) bool {
	for _, attachment := range ig.Attachments {
		if attachment.VpcId != nil && *attachment.VpcId == *vpc.VpcId {
			return true
		}
	}
	return false
}

func attachInternetGatewayToVpc(ig *ec2.InternetGateway, vpc *ec2.Vpc) {
	_, err := amazon.EC2().AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		InternetGatewayId: ig.InternetGatewayId,
		VpcId:             vpc.VpcId,
	})
	util.Check(err)
}

func detachInternetGatewayFromVpc(ig *ec2.InternetGateway, vpc *ec2.Vpc) {
	_, err := amazon.EC2().DetachInternetGateway(&ec2.DetachInternetGatewayInput{
		InternetGatewayId: ig.InternetGatewayId,
		VpcId:             vpc.VpcId,
	})
	util.Check(err)
}
