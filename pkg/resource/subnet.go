package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func SelectSubnet(lr *LabResources, isPrivateSubnet bool) string {
	if isPrivateSubnet {
		return *lr.PrivateSubnet.SubnetId
	}
	return *lr.PublicSubnet.SubnetId
}

func resolvePublicSubnetAttributes(publicSubnet *ec2.Subnet) {
	if !*publicSubnet.MapPublicIpOnLaunch {
		_, err := amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
			SubnetId:            publicSubnet.SubnetId,
			MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
		})
		util.Check(err)
	}
}
