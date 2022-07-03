package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (co *AWSCloudOperator) SubnetId(isPrivateSubnet bool) string {
	if isPrivateSubnet {
		return *co.Rs.PrivateSubnet.SubnetId
	}
	return *co.Rs.PublicSubnet.SubnetId
}

func (co *AWSCloudOperator) resolvePublicSubnetAttributes() {
	if !*co.Rs.PublicSubnet.MapPublicIpOnLaunch {
		_, err := amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
			SubnetId:            co.Rs.PublicSubnet.SubnetId,
			MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
		})
		util.MustExec(err)
	}
}
