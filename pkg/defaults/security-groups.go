package defaults

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (cldo *CloudLabDefaultsOperator) findCloudLabSecurityGroups() {
	cldo.SecurityGroups = []*ec2.SecurityGroup{}
	err := amazon.EC2().DescribeSecurityGroupsPages(
		&ec2.DescribeSecurityGroupsInput{},
		func(dsgo *ec2.DescribeSecurityGroupsOutput, b bool) bool {
			for _, sg := range dsgo.SecurityGroups {
				if NameTagEquals(sg.Tags, CloudLabSecutiyGroup) {
					cldo.SecurityGroups = append(cldo.SecurityGroups, sg)
					continue
				}
				if cldo.isVpcDefaultSecurityGroup(sg.VpcId) {
					cldo.nameDefaultSecutiyGroup(sg.GroupId)
					cldo.SecurityGroups = append(cldo.SecurityGroups, sg)
					continue
				}
			}
			return true
		},
	)
	util.MustExec(err)
}

func (cldo *CloudLabDefaultsOperator) isVpcDefaultSecurityGroup(vpcId *string) bool {
	if cldo.Vpc != nil && *cldo.Vpc.VpcId == *vpcId {
		return true
	}
	return false
}

func (cldo *CloudLabDefaultsOperator) nameDefaultSecutiyGroup(securityGroupId *string) {
	_, err := amazon.EC2().CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{securityGroupId},
		Tags:      createNameTagArray(CloudLabSecutiyGroup),
	})
	util.MustExec(err)
}

func (cldo *CloudLabDefaultsOperator) CreateSecurityGroup(name string, port int) {
	for _, sg := range cldo.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == name {
			return
		}
	}
	csgo, err := amazon.EC2().CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		VpcId:             cldo.Vpc.VpcId,
		GroupName:         util.StrPtr(name),
		Description:       util.StrPtr(name),
		TagSpecifications: CreateNameTagSpec("security-group", CloudLabSecutiyGroup),
	})
	util.MustExec(err)
	createInboundRule(csgo.GroupId, "tcp", port)
	createInboundRule(csgo.GroupId, "udp", port)
	util.VMessage("created", name, *csgo.GroupId)
	cldo.findCloudLabSecurityGroups()
}

func createInboundRule(groupId *string, protocol Protocol, port int) {
	_, err := amazon.EC2().AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:    groupId,
		FromPort:   util.IntToInt64Ptr(port),
		ToPort:     util.IntToInt64Ptr(port),
		IpProtocol: util.StrPtr(string(protocol)),
		CidrIp:     util.StrPtr(AllIpsCidr),
	})
	util.MustExec(err)
}

func (cldo *CloudLabDefaultsOperator) deleteAllSecurityGroups() {
	for _, sg := range cldo.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == "default" {
			continue
		}
		_, err := amazon.EC2().DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
			GroupId: sg.GroupId,
		})
		util.MustExec(err)
		util.VMessage("deleted", CloudLabSecutiyGroup, *sg.GroupId)
	}
}

func (cldo *CloudLabDefaultsOperator) Port22() *string {
	for _, sg := range cldo.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == "22" {
			return sg.GroupId
		}
	}
	cldo.CreateSecurityGroup("22", 22)
	cldo.findCloudLabSecurityGroups()
	return cldo.GetSecurityGroupIdByNameOrPanic("22").GroupId
}

func (cldo *CloudLabDefaultsOperator) GetSecurityGroupIdByNameOrPanic(name string) *ec2.SecurityGroup {
	for _, sg := range cldo.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	panic(fmt.Sprintf("failed to find security group %s", name))
}

func (cldo *CloudLabDefaultsOperator) GetSecurityGroupIdByNameOrNil(name string) *ec2.SecurityGroup {
	for _, sg := range cldo.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	return nil
}
