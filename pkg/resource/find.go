package resource

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func findVpc(name string) (vpcToFind *ec2.Vpc) {
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

func findInstances() (nodes []*ec2.Instance) {
	err := amazon.EC2().DescribeInstancesPages(&ec2.DescribeInstancesInput{},
		func(dio *ec2.DescribeInstancesOutput, b bool) bool {
			for _, res := range dio.Reservations {
				for _, inst := range res.Instances {
					if NameTagEquals(inst.Tags, CloudLabInstance) {
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

func findNotTerminatedInstances(nodes []*ec2.Instance) (notTerminatedNodes []*ec2.Instance) {
	for _, node := range nodes {
		if *node.State.Name != "terminated" {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (ro *ResourceOperator) FindInstanceByName(name string) *ec2.Instance {
	for _, inst := range ro.Instances {
		if NameTagEquals(inst.Tags, name) {
			return inst
		}
	}
	panic("node not found")
}

func findInternetGateway() (igToFind *ec2.InternetGateway) {
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

func FindAllCloudLabKeyPairs() (kps []*ec2.KeyPairInfo) {
	dkpo, err := amazon.EC2().DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	util.MustExec(err)

	for _, kp := range dkpo.KeyPairs {
		if NameTagEquals(kp.Tags, CloudLabKeyPair) {
			kps = append(kps, kp)
		}
	}

	return kps
}

func FindCurrentKeyPair(kps []*ec2.KeyPairInfo) *ec2.KeyPairInfo {
	if kps == nil {
		return nil
	}
	currentKeyPairName := GetCurrentCloudLabKeyPairName(kps)
	for _, kp := range kps {
		if kp.KeyName != nil && *kp.KeyName == currentKeyPairName {
			return kp
		}
	}
	return nil
}

func findRouteTable(vpcId string, name string) (routeTableToFind *ec2.RouteTable) {
	err := amazon.EC2().DescribeRouteTablesPages(
		&ec2.DescribeRouteTablesInput{},
		func(drto *ec2.DescribeRouteTablesOutput, b bool) bool {
			for _, rt := range drto.RouteTables {
				if rt.VpcId != nil && *rt.VpcId == vpcId {
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

func findMainRouteTable(vpc *ec2.Vpc) (routeTableToFind *ec2.RouteTable) {
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

func findCloudLabSecurityGroups() (sgs []*ec2.SecurityGroup) {
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

func findSecurityGroupByName(sgs []*ec2.SecurityGroup, name string) *ec2.SecurityGroup {
	for _, sg := range sgs {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	return nil
}

func findSubnet(name string) (subnetToFind *ec2.Subnet) {
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
