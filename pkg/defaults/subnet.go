package defaults

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (cldo *CloudLabDefaultsOperator) createSubnets() {
	if !cldo.foundPublicSubnet() {
		cldo.PublicSubnet = cldo.createSubnet(CloudLabPublicSubnetName, PublicSubnetCidrBlock)
	}
	if !cldo.foundPrivateSubnet() {
		cldo.PrivateSubnet = cldo.createSubnet(CloudLabPrivateSubnetName, PrivateSubnetCidrBlock)
	}
	cldo.resolvePublicSubnetAttributes()
}

func (cldo *CloudLabDefaultsOperator) deleteSubnets() {
	if cldo.foundPublicSubnet() {
		deleteSubnet(CloudLabPublicSubnetName, cldo.PublicSubnet.SubnetId)
	}
	if cldo.foundPrivateSubnet() {
		deleteSubnet(CloudLabPrivateSubnetName, cldo.PrivateSubnet.SubnetId)
	}
}

func (cldo *CloudLabDefaultsOperator) foundPublicSubnet() bool {
	return cldo.PublicSubnet != nil
}

func (cldo *CloudLabDefaultsOperator) foundPrivateSubnet() bool {
	return cldo.PrivateSubnet != nil
}

func findSubnet(name string) (sn *ec2.Subnet) {
	err := amazon.EC2().DescribeSubnetsPages(
		&ec2.DescribeSubnetsInput{},
		func(dso *ec2.DescribeSubnetsOutput, b bool) bool {
			for _, subnet := range dso.Subnets {
				nameTagValue := findNameTagValue(subnet.Tags)
				if nameTagValue != nil && *nameTagValue == name {
					sn = subnet
					return false
				}
			}
			return true
		})
	util.MustExec(err)
	return sn
}

func (cldo CloudLabDefaultsOperator) createSubnet(name string, cidr string) (sn *ec2.Subnet) {
	cso, err := amazon.EC2().CreateSubnet(&ec2.CreateSubnetInput{
		VpcId:             cldo.Vpc.VpcId,
		CidrBlock:         util.StrPtr(cidr),
		TagSpecifications: CreateNameTagSpec("subnet", name),
	})
	util.MustExec(err)
	util.VMessage("created", name, *cso.Subnet.SubnetId)
	return cso.Subnet
}

func deleteSubnet(name string, subnetId *string) {
	_, err := amazon.EC2().DeleteSubnet(&ec2.DeleteSubnetInput{
		SubnetId: subnetId,
	})
	util.MustExec(err)
	util.VMessage("deleted", name, *subnetId)
}

func (cldo CloudLabDefaultsOperator) resolvePublicSubnetAttributes() {
	if !*cldo.PublicSubnet.MapPublicIpOnLaunch {
		_, err := amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
			SubnetId:            cldo.PublicSubnet.SubnetId,
			MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
		})
		util.MustExec(err)
	}
	// if !*cldo.PublicSubnet.PrivateDnsNameOptionsOnLaunch.EnableResourceNameDnsARecord {
	// 	_, err := amazon.EC2().ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
	// 		SubnetId:                             cldo.PublicSubnet.SubnetId,
	// 		EnableResourceNameDnsARecordOnLaunch: &ec2.AttributeBooleanValue{Value: util.BoolPtr(true)},
	// 	})
	// 	util.MustExec(err)
	// }
}
