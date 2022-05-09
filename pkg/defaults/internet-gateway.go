package defaults

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (cldo *CloudLabDefaultsOperator) findInternetGateway() {
	err := amazon.EC2().DescribeInternetGatewaysPages(
		&ec2.DescribeInternetGatewaysInput{},
		func(digo *ec2.DescribeInternetGatewaysOutput, b bool) bool {
			for _, ig := range digo.InternetGateways {
				if nameTagEquals(ig.Tags, CloudLabInternetGateway) {
					cldo.InternetGateway = ig
					return false
				}
			}
			return true
		},
	)
	util.MustExec(err)
}

func (cldo *CloudLabDefaultsOperator) internetGatewayExists() bool {
	return cldo.InternetGateway != nil
}

func (cldo *CloudLabDefaultsOperator) createInternetGateway() {
	if cldo.internetGatewayExists() {
		return
	}
	cigo, err := amazon.EC2().CreateInternetGateway(&ec2.CreateInternetGatewayInput{
		TagSpecifications: CreateNameTagSpec("internet-gateway", CloudLabInternetGateway),
	})
	util.MustExec(err)
	cldo.InternetGateway = cigo.InternetGateway
	util.VMessage("created", CloudLabInternetGateway, *cigo.InternetGateway.InternetGatewayId)
}

func (cldo *CloudLabDefaultsOperator) deleteInternetGateway() {
	if !cldo.internetGatewayExists() {
		return
	}
	_, err := amazon.EC2().DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
		InternetGatewayId: cldo.InternetGateway.InternetGatewayId,
	})
	util.MustExec(err)
	util.VMessage("deleted", CloudLabInternetGateway, *cldo.InternetGateway.InternetGatewayId)
}

func (cldo *CloudLabDefaultsOperator) isAttachedToVpc() bool {
	if cldo.InternetGateway == nil {
		return false
	}
	for _, attachment := range cldo.InternetGateway.Attachments {
		if attachment.VpcId != nil && *attachment.VpcId == *cldo.Vpc.VpcId {
			return true
		}
	}
	return false
}

func (cldo *CloudLabDefaultsOperator) attachInternetGateway() {
	if cldo.isAttachedToVpc() {
		return
	}
	_, err := amazon.EC2().AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		InternetGatewayId: cldo.InternetGateway.InternetGatewayId,
		VpcId:             cldo.Vpc.VpcId,
	})
	util.MustExec(err)
	fmt.Println(fmt.Sprintf("%s is attached to %s", CloudLabInternetGateway, DefaultVpcName))
}

func (cldo *CloudLabDefaultsOperator) detachInternetGateway() {
	if !cldo.isAttachedToVpc() {
		return
	}
	_, err := amazon.EC2().DetachInternetGateway(&ec2.DetachInternetGatewayInput{
		InternetGatewayId: cldo.InternetGateway.InternetGatewayId,
		VpcId:             cldo.Vpc.VpcId,
	})
	util.MustExec(err)
	fmt.Println(fmt.Sprintf("%s is detached from %s", CloudLabInternetGateway, DefaultVpcName))
}
