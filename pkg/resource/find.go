package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// #############################################################################
// #############################################################################
// #############################################################################

func FindVpc() (targetVpc *ec2.Vpc) {
	err := amazon.EC2().DescribeVpcsPages(
		&ec2.DescribeVpcsInput{},
		func(dvo *ec2.DescribeVpcsOutput, b bool) bool {
			for _, vpc := range dvo.Vpcs {
				nameTagValue := FindNameTagValue(vpc.Tags)
				if nameTagValue != nil && *nameTagValue == CloudLabVpc {
					targetVpc = vpc
					return false
				}
			}
			return true
		})
	util.Check(err)
	util.Log("found %v: %v", CloudLabVpc, targetVpc != nil)
	return targetVpc
}

func FindVpcOrPanic() (targetVpc *ec2.Vpc) {
	targetVpc = FindVpc()
	if targetVpc == nil {
		panic("run 'lab init' first")
	}
	util.Log("found %v: %v", CloudLabVpc, targetVpc != nil)
	return targetVpc
}

// #############################################################################
// #############################################################################
// #############################################################################

func FindInstances() (instances []*ec2.Instance) {
	err := amazon.EC2().DescribeInstancesPages(&ec2.DescribeInstancesInput{},
		func(dio *ec2.DescribeInstancesOutput, b bool) bool {
			for _, res := range dio.Reservations {
				for _, inst := range res.Instances {
					if TagEquals(inst.Tags, isCloudLabInstanceTagKey, isCloudLabInstanceTagVal) {
						instances = append(instances, inst)
					}
				}
			}
			return true
		},
	)
	util.Check(err)
	util.Log("found %v cl instances", len(instances))
	return instances
}

func FindNonTerminatedInstances() (instances []*ec2.Instance) {
	err := amazon.EC2().DescribeInstancesPages(&ec2.DescribeInstancesInput{},
		func(dio *ec2.DescribeInstancesOutput, b bool) bool {
			for _, res := range dio.Reservations {
				for _, inst := range res.Instances {
					if TagEquals(inst.Tags, isCloudLabInstanceTagKey, isCloudLabInstanceTagVal) {
						if inst.State.Name != nil && *inst.State.Name != "terminated" {
							instances = append(instances, inst)
						}
					}
				}
			}
			return true
		},
	)
	util.Check(err)
	util.Log("found %v cl instances", len(instances))
	return instances
}

func FindInstanceByName(name string) (instance *ec2.Instance) {
	err := amazon.EC2().DescribeInstancesPages(&ec2.DescribeInstancesInput{},
		func(dio *ec2.DescribeInstancesOutput, b bool) bool {
			for _, res := range dio.Reservations {
				for _, inst := range res.Instances {
					if NameTagEquals(inst.Tags, name) {
						instance = inst
						return false
					}
				}
			}
			return true
		},
	)
	util.Check(err)
	util.Log("found instance %v: %v", name, instance != nil)
	return instance
}

func FindInstanceByNameOrPanic(name string) (instance *ec2.Instance) {
	instance = FindInstanceByName(name)
	if instance == nil {
		panic(fmt.Sprintf("instance '%s' not found", name))
	}
	util.Log("found instance %v: %v", name, instance != nil)
	return instance
}

// #############################################################################
// #############################################################################
// #############################################################################

func findInternetGateway(name string) (targetIg *ec2.InternetGateway) {
	err := amazon.EC2().DescribeInternetGatewaysPages(
		&ec2.DescribeInternetGatewaysInput{},
		func(digo *ec2.DescribeInternetGatewaysOutput, b bool) bool {
			for _, ig := range digo.InternetGateways {
				if NameTagEquals(ig.Tags, name) {
					targetIg = ig
					return false
				}
			}
			return true
		},
	)
	util.Check(err)
	util.Log("found internet gateway %v: %v", name, targetIg != nil)
	return targetIg
}

// #############################################################################
// #############################################################################
// #############################################################################

func findKeyPair() (keypair *ec2.KeyPairInfo) {
	dkpo, err := amazon.EC2().DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	util.Check(err)
	for _, kp := range dkpo.KeyPairs {
		if NameTagEquals(kp.Tags, CloudLabKeyPair) {
			keypair = kp
		}
	}
	util.Log("found cl keypair: %v", keypair != nil)
	return keypair
}

// #############################################################################
// #############################################################################
// #############################################################################

func findRouteTable(vpc *ec2.Vpc, name string) (targetRT *ec2.RouteTable) {
	err := amazon.EC2().DescribeRouteTablesPages(
		&ec2.DescribeRouteTablesInput{},
		func(drto *ec2.DescribeRouteTablesOutput, b bool) bool {
			for _, rt := range drto.RouteTables {
				if rt.VpcId != nil && *rt.VpcId == *vpc.VpcId {
					if NameTagEquals(rt.Tags, name) {
						targetRT = rt
						return false
					}
				}
			}
			return true
		},
	)
	util.Check(err)
	util.Log("found route table %v: %v", name, targetRT == nil)
	return targetRT
}

func findMainRouteTable(vpc *ec2.Vpc) (targetRT *ec2.RouteTable) {
	err := amazon.EC2().DescribeRouteTablesPages(
		&ec2.DescribeRouteTablesInput{},
		func(drto *ec2.DescribeRouteTablesOutput, b bool) bool {
			for _, rt := range drto.RouteTables {
				if rt.VpcId != nil && *rt.VpcId == *vpc.VpcId {
					for _, assoc := range rt.Associations {
						if assoc.Main != nil && *assoc.Main {
							targetRT = rt
							return false
						}
					}
				}
			}
			return true
		},
	)
	util.Check(err)
	util.Log("found main route table: %v", targetRT != nil)
	return targetRT
}

// #############################################################################
// #############################################################################
// #############################################################################

func findSecurityGroupByName(sgs []*ec2.SecurityGroup, name string) *ec2.SecurityGroup {
	for _, sg := range sgs {
		if sg.GroupName != nil && *sg.GroupName == name {
			return sg
		}
	}
	return nil
}

func FindAllSecurityGroups() (sgs []*ec2.SecurityGroup) {
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
	util.Check(err)
	util.Log("found %v cl security groups", len(sgs))
	return sgs
}

// #############################################################################
// #############################################################################
// #############################################################################

func findSubnet(name string) (targetSubnet *ec2.Subnet) {
	err := amazon.EC2().DescribeSubnetsPages(
		&ec2.DescribeSubnetsInput{},
		func(dso *ec2.DescribeSubnetsOutput, b bool) bool {
			for _, subnet := range dso.Subnets {
				nameTagValue := FindNameTagValue(subnet.Tags)
				if nameTagValue != nil && *nameTagValue == name {
					targetSubnet = subnet
					return false
				}
			}
			return true
		})
	util.Check(err)
	util.Log("found subnet %v: %v", name, targetSubnet != nil)
	return targetSubnet
}

func FindPublicSubnet() *ec2.Subnet {
	return findSubnet(CloudLabPublicSubnet)
}
func FindPrivateSubnet() *ec2.Subnet {
	return findSubnet(CloudLabPrivateSubnet)
}

// #############################################################################
// #############################################################################
// #############################################################################
