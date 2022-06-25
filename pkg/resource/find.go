package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type ResourceFinder interface {
	findVpc(name string) *ec2.Vpc
	findInstances() []*ec2.Instance
	findNotTerminatedInstances(nodes []*ec2.Instance) []*ec2.Instance
	findInternetGateway() *ec2.InternetGateway
	findAllCloudLabKeyPairs() []*ec2.KeyPairInfo
	findKeyPair() *ec2.KeyPairInfo
	findRouteTable(vpc *ec2.Vpc, name string) *ec2.RouteTable
	findMainRouteTable(vpc *ec2.Vpc) *ec2.RouteTable
	findCloudLabSecurityGroups() []*ec2.SecurityGroup
	findSecurityGroupByName(sgs []*ec2.SecurityGroup, name string) *ec2.SecurityGroup
	findSubnet(name string) *ec2.Subnet
}

type AWSResourceFinder struct{}

func (a *AWSResourceFinder) findVpc(name string) (vpcToFind *ec2.Vpc) {
	err := amazon.EC2().DescribeVpcsPages(
		&ec2.DescribeVpcsInput{},
		func(dvo *ec2.DescribeVpcsOutput, b bool) bool {
			for _, vpc := range dvo.Vpcs {
				nameTagValue := FindNameTagValue(vpc.Tags)
				if nameTagValue != nil && *nameTagValue == name {
					vpcToFind = vpc
					return false
				}
			}
			return true
		})
	util.MustExec(err)
	return vpcToFind
}

func (a *AWSResourceFinder) findInstances() (nodes []*ec2.Instance) {
	err := amazon.EC2().DescribeInstancesPages(&ec2.DescribeInstancesInput{},
		func(dio *ec2.DescribeInstancesOutput, b bool) bool {
			for _, res := range dio.Reservations {
				for _, inst := range res.Instances {
					if TagEquals(inst.Tags, IsCloudLabInstanceTagKey, IsCloudLabInstanceTagVal) {
						nodes = append(nodes, inst)
					}
				}
			}
			return true
		},
	)
	util.MustExec(err)
	return nodes
}

func (a *AWSResourceFinder) findNotTerminatedInstances(nodes []*ec2.Instance) (notTerminatedNodes []*ec2.Instance) {
	for _, node := range nodes {
		if node.State != nil && node.State.Name != nil && *node.State.Name != "terminated" {
			nodes = append(nodes, node)
		}
	}
	return notTerminatedNodes
}

func (a *AWSResourceFinder) findInternetGateway() (igToFind *ec2.InternetGateway) {
	err := amazon.EC2().DescribeInternetGatewaysPages(
		&ec2.DescribeInternetGatewaysInput{},
		func(digo *ec2.DescribeInternetGatewaysOutput, b bool) bool {
			for _, ig := range digo.InternetGateways {
				if NameTagEquals(ig.Tags, CloudLabInternetGateway) {
					igToFind = ig
					return false
				}
			}
			return true
		},
	)
	util.MustExec(err)
	return igToFind
}

func (a *AWSResourceFinder) findAllCloudLabKeyPairs() (kps []*ec2.KeyPairInfo) {
	dkpo, err := amazon.EC2().DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	util.MustExec(err)

	for _, kp := range dkpo.KeyPairs {
		if NameTagEquals(kp.Tags, CloudLabKeyPair) {
			kps = append(kps, kp)
		}
	}

	return kps
}

func (a *AWSResourceFinder) findKeyPair() (keypair *ec2.KeyPairInfo) {
	dkpo, err := amazon.EC2().DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	util.MustExec(err)

	for _, kp := range dkpo.KeyPairs {
		if NameTagEquals(kp.Tags, CloudLabKeyPair) {
			keypair = kp
		}
	}

	return keypair
}

func (a *AWSResourceFinder) findRouteTable(vpc *ec2.Vpc, name string) (routeTableToFind *ec2.RouteTable) {
	err := amazon.EC2().DescribeRouteTablesPages(
		&ec2.DescribeRouteTablesInput{},
		func(drto *ec2.DescribeRouteTablesOutput, b bool) bool {
			for _, rt := range drto.RouteTables {
				if rt.VpcId != nil && *rt.VpcId == *vpc.VpcId {
					if NameTagEquals(rt.Tags, name) {
						routeTableToFind = rt
						return false
					}
				}
			}
			return true
		},
	)
	util.MustExec(err)
	return routeTableToFind
}

func (a *AWSResourceFinder) findMainRouteTable(vpc *ec2.Vpc) (routeTableToFind *ec2.RouteTable) {
	err := amazon.EC2().DescribeRouteTablesPages(
		&ec2.DescribeRouteTablesInput{},
		func(drto *ec2.DescribeRouteTablesOutput, b bool) bool {
			for _, rt := range drto.RouteTables {
				if rt.VpcId != nil && *rt.VpcId == *vpc.VpcId {
					for _, assoc := range rt.Associations {
						if *assoc.Main {
							routeTableToFind = rt
							return false
						}
					}
				}
			}
			return true
		},
	)
	util.MustExec(err)
	return routeTableToFind
}

func (a *AWSResourceFinder) findCloudLabSecurityGroups() (sgs []*ec2.SecurityGroup) {
	err := amazon.EC2().DescribeSecurityGroupsPages(
		&ec2.DescribeSecurityGroupsInput{},
		func(dsgo *ec2.DescribeSecurityGroupsOutput, b bool) bool {
			for _, sg := range dsgo.SecurityGroups {
				if NameTagEquals(sg.Tags, CloudLabSecutiyGroup) {
					sgs = append(sgs, sg)
					continue
				}
			}
			return true
		},
	)
	util.MustExec(err)
	return sgs
}

func (a *AWSResourceFinder) findSecurityGroupByName(sgs []*ec2.SecurityGroup, name string) *ec2.SecurityGroup {
	for _, sg := range sgs {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	return nil
}

func (a *AWSResourceFinder) findSubnet(name string) (subnetToFind *ec2.Subnet) {
	err := amazon.EC2().DescribeSubnetsPages(
		&ec2.DescribeSubnetsInput{},
		func(dso *ec2.DescribeSubnetsOutput, b bool) bool {
			for _, subnet := range dso.Subnets {
				nameTagValue := FindNameTagValue(subnet.Tags)
				if nameTagValue != nil && *nameTagValue == name {
					subnetToFind = subnet
					return false
				}
			}
			return true
		})
	util.MustExec(err)
	return subnetToFind
}
