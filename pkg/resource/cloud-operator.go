package resource

import (
	"fmt"
	"os"
)

// ######################################################################################
// ######################################################################################
// ######################################################################################

type AWSCloudOperator struct {
	creator ResourceCreator
	finder  ResourceFinder
	deleter ResourceDeleter
	Rs      *AWSResources
}

func new() *AWSCloudOperator {
	return &AWSCloudOperator{
		creator: &AWSCreator{},
		finder:  &AWSResourceFinder{},
		Rs:      &AWSResources{},
	}
}

func NewCloudOperatorNoAudit() *AWSCloudOperator {
	co := new()
	co.FindAll()
	return co
}

// Finds all resources and audits them.
func NewCloudOperator() *AWSCloudOperator {
	co := new()
	co.FindAll()
	co.Audit()
	return co
}

// ######################################################################################
// ######################################################################################
// ######################################################################################

func (co *AWSCloudOperator) FindAll() {
	co.Rs.Vpc = co.finder.findVpc(CloudLabVpc)
	if co.Rs.Vpc == nil {
		return
	}
	co.Rs.PublicSubnet = co.finder.findSubnet(CloudLabPublicSubnet)
	co.Rs.PrivateSubnet = co.finder.findSubnet(CloudLabPrivateSubnet)
	co.Rs.PublicRouteTable = co.finder.findMainRouteTable(co.Rs.Vpc)
	co.Rs.PrivateRouteTable = co.finder.findRouteTable(co.Rs.Vpc, CloudLabPrivateRouteTable)
	co.Rs.InternetGateway = co.finder.findInternetGateway()
	co.Rs.SecurityGroups = co.finder.findCloudLabSecurityGroups()
	co.Rs.Instances = co.finder.findInstances()
	co.Rs.KeyPair = co.finder.findKeyPair()
}

func (co *AWSCloudOperator) Audit() {
	FatalMissing(co.Rs.Vpc == nil, "vpc")
	FatalMissing(co.Rs.PublicSubnet == nil, "public subnet")
	FatalMissing(co.Rs.PrivateSubnet == nil, "private subnet")
	FatalMissing(co.Rs.PublicRouteTable == nil, "main route table")
	FatalMissing(co.Rs.PrivateRouteTable == nil, "private route table")
	FatalMissing(co.Rs.InternetGateway == nil, "internet gateway")
	FatalMissing(co.Rs.SecurityGroups == nil, "security groups")
	FatalMissing(co.Rs.KeyPair == nil, "key pair")
}

func FatalMissing(missing bool, resourceType string) {
	if missing {
		fmt.Printf("%s is missing from cloudlab resources\n", resourceType)
		fmt.Println("run 'lab init' to create missing cloudlab resources")
		os.Exit(1)
	}
}

// ######################################################################################
// Init #################################################################################
// ######################################################################################

func (co *AWSCloudOperator) InitializeCloudLabResources() {
	if co.Rs.Vpc == nil {
		co.Rs.Vpc = co.creator.createVpc(DefaultVpcCidrBlock, CloudLabVpc)
		co.Rs.PublicRouteTable = co.finder.findMainRouteTable(co.Rs.Vpc)
	}
	if co.Rs.PublicSubnet == nil {
		co.Rs.PublicSubnet = co.creator.createSubnet(co.Rs.Vpc, CloudLabPublicSubnet, PublicSubnetCidrBlock)
		co.resolvePublicSubnetAttributes()
	}
	if co.Rs.PrivateSubnet == nil {
		co.Rs.PrivateSubnet = co.creator.createSubnet(co.Rs.Vpc, CloudLabPrivateSubnet, PrivateSubnetCidrBlock)
	}
	if co.Rs.PrivateRouteTable == nil {
		co.Rs.PrivateRouteTable = co.creator.createRouteTable(co.Rs.Vpc, CloudLabPrivateRouteTable)
	}
	if co.Rs.InternetGateway == nil {
		co.Rs.InternetGateway = co.creator.createInternetGateway(CloudLabInternetGateway)
	}
	if !InternetGatewayIsAttachedToVpc(co.Rs.InternetGateway, co.Rs.Vpc) {
		attachInternetGatewayToVpc(co.Rs.InternetGateway, co.Rs.Vpc)
	}
	if !internetGatewayRouteExistsOnRouteTable(co.Rs.PublicRouteTable, co.Rs.InternetGateway) {
		addInternetGatewayRoute(co.Rs.PublicRouteTable, co.Rs.InternetGateway, RouteTablePublicSubnetCidr)
	}
	if !subnetAssociationExistsOnRouteTable(co.Rs.PublicRouteTable, co.Rs.PublicSubnet) {
		associateSubnetWithRouteTable(co.Rs.PublicRouteTable, co.Rs.PublicSubnet)
	}
	securityGroup22 := co.finder.findSecurityGroupByName(co.Rs.SecurityGroups, "22")
	if securityGroup22 == nil {
		CreateSecurityGroup(co.Rs.Vpc, "22", 22)
		co.Rs.SecurityGroups = co.finder.findCloudLabSecurityGroups()
	}
	if co.Rs.KeyPair == nil {
		CreateKeyPair()
		co.Rs.KeyPair = co.finder.findKeyPair()
	}
}

// ######################################################################################
// ######################################################################################
// ######################################################################################

func (co *AWSCloudOperator) DestroyCloudLabResources() {
	if len(co.finder.findNotTerminatedInstances(co.Rs.Instances)) > 0 {
		panic("run 'lab delete all' and try again")
	}

	if len(co.Rs.Instances) > 0 {
		co.deleter.deleteInstances(co.Rs.Instances)
	}
	if co.Rs.InternetGateway != nil {
		detachInternetGatewayFromVpc(co.Rs.InternetGateway, co.Rs.Vpc)
		co.deleter.deleteInternetGateway(co.Rs.InternetGateway)
	}

	if co.Rs.PrivateRouteTable != nil {
		disassociateSubnetsFromRouteTable(co.Rs.PrivateRouteTable)
		co.deleter.deleteRouteTable(co.Rs.PrivateRouteTable)
	}

	if co.Rs.PublicSubnet != nil {
		co.deleter.deleteSubnet(co.Rs.PublicSubnet)
	}
	if co.Rs.PrivateSubnet != nil {
		co.deleter.deleteSubnet(co.Rs.PrivateSubnet)
	}

	co.deleter.deleteSecurityGroups(co.Rs.SecurityGroups)

	if co.Rs.Vpc != nil {
		co.deleter.deleteVpc(co.Rs.Vpc)
	}

	co.deleter.deleteKeyPair(co.Rs.KeyPair)
}
