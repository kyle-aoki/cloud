package resource

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func internetGatewayIsAttachedToVpc(ig *ec2.InternetGateway, vpc *ec2.Vpc) bool {
	for _, attachment := range ig.Attachments {
		if attachment.VpcId != nil && *attachment.VpcId == *vpc.VpcId {
			return true
		}
	}
	return false
}

func attachInternetGateway(ig *ec2.InternetGateway, vpc *ec2.Vpc) {
	_, err := amazon.EC2().AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		InternetGatewayId: ig.InternetGatewayId,
		VpcId:             vpc.VpcId,
	})
	util.MustExec(err)
}

func detachInternetGateway(ig *ec2.InternetGateway, vpc *ec2.Vpc) {
	_, err := amazon.EC2().DetachInternetGateway(&ec2.DetachInternetGatewayInput{
		InternetGatewayId: ig.InternetGatewayId,
		VpcId:             vpc.VpcId,
	})
	util.MustExec(err)
}
