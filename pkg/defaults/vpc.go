package defaults

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (cldo *CloudLabDefaultsOperator) createVpc() {
	if !cldo.foundVpc() {
		cldo.createCloudLabVpc()
	}
}

func (cldo *CloudLabDefaultsOperator) deleteVpc() {
	if cldo.foundVpc() {
		cldo.deleteCloudLabVpc()
	}
}

func (cldo *CloudLabDefaultsOperator) foundVpc() bool {
	return cldo.Vpc != nil
}

func (cldo *CloudLabDefaultsOperator) findCloudLabVpc() {
	err := amazon.EC2().DescribeVpcsPages(
		&ec2.DescribeVpcsInput{},
		func(dvo *ec2.DescribeVpcsOutput, b bool) bool {
			for _, vpc := range dvo.Vpcs {
				nameTagValue := findNameTagValue(vpc.Tags)
				if nameTagValue != nil && *nameTagValue == DefaultVpcName {
					cldo.Vpc = vpc
					return false
				}
			}
			return true
		})
	util.MustExec(err)
}

func (cldo *CloudLabDefaultsOperator) createCloudLabVpc() {
	cvo, err := amazon.EC2().CreateVpc(&ec2.CreateVpcInput{
		CidrBlock:         util.StrPtr(DefaultVpcCidrBlock),
		TagSpecifications: createNameTag("vpc", DefaultVpcName),
	})
	util.MustExec(err)
	cldo.Vpc = cvo.Vpc
	util.VMessage("created", DefaultVpcName, *cvo.Vpc.VpcId)
}

func (cldo *CloudLabDefaultsOperator) deleteCloudLabVpc() {
	_, err := amazon.EC2().DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: cldo.Vpc.VpcId,
	})
	util.MustExec(err)
	util.VMessage("deleted", DefaultVpcName, *cldo.Vpc.VpcId)
}
