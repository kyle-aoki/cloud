package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (ro *ResourceOperator) resolvePublicSubnetAttributes() {
	if !*ro.PublicSubnet.MapPublicIpOnLaunch {
		_, err := amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
			SubnetId:            ro.PublicSubnet.SubnetId,
			MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
		})
		util.MustExec(err)
	}
}
