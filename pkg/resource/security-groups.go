package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func assignNameTagToDefaultSecurityGroupIfMissing(cloudlabVpcId string) {
	err := amazon.EC2().DescribeSecurityGroupsPages(
		&ec2.DescribeSecurityGroupsInput{},
		func(dsgo *ec2.DescribeSecurityGroupsOutput, b bool) bool {
			for _, sg := range dsgo.SecurityGroups {
				if cloudlabVpcId == *sg.VpcId {
					if !NameTagEquals(sg.Tags, CloudLabSecutiyGroup) {
						nameDefaultSecutiyGroup(sg.GroupId, CloudLabSecutiyGroup)
					}
					return false
				}
			}
			return true
		},
	)
	util.MustExec(err)
}

// assign name tag to default sg if name tag is missing
func nameDefaultSecutiyGroup(securityGroupId *string, name string) {
	_, err := amazon.EC2().CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{securityGroupId},
		Tags:      CreateNameTagArray(name),
	})
	util.MustExec(err)
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

func (co *AWSCloudOperator) GetSecurityGroupIdByNameOrPanic(name string) *ec2.SecurityGroup {
	for _, sg := range co.Rs.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	panic(fmt.Sprintf("failed to find security group %s", name))
}

func (co *AWSCloudOperator) GetSecurityGroupIdByNameOrNil(name string) *ec2.SecurityGroup {
	for _, sg := range co.Rs.SecurityGroups {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	return nil
}