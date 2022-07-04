package resource

import (
	"cloudlab/pkg/util"
	"fmt"
	"log"
	"os"
)

// ######################################################################################
// ######################################################################################
// ######################################################################################

type AWSCloudOperator struct {
	Creator *ResourceCreator
	Finder  *ResourceFinder
	Deleter *ResourceDeleter
	Rs      *AWSResources
}

func new() *AWSCloudOperator {
	return &AWSCloudOperator{
		Creator: &ResourceCreator{},
		Finder:  &ResourceFinder{},
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
	log.Println("finding cloudlab resources...")
	co.Rs.Vpc = co.Finder.findVpc(CloudLabVpc)
	if co.Rs.Vpc == nil {
		return
	}
	co.Rs.PublicSubnet = co.Finder.findSubnet(CloudLabPublicSubnet)
	co.Rs.PrivateSubnet = co.Finder.findSubnet(CloudLabPrivateSubnet)
	co.Rs.PublicRouteTable = co.Finder.findMainRouteTable(co.Rs.Vpc)
	co.Rs.PrivateRouteTable = co.Finder.findRouteTable(co.Rs.Vpc, CloudLabPrivateRouteTable)
	co.Rs.InternetGateway = co.Finder.findInternetGateway(CloudLabInternetGateway)
	co.Rs.SecurityGroups = co.Finder.FindSecurityGroups()
	co.Rs.Instances = co.Finder.findInstances()
	co.Rs.KeyPair = co.Finder.findKeyPair()
}

func (co *AWSCloudOperator) Audit() {
	bools := []bool{
		IsMissing(co.Rs.Vpc == nil, "vpc"),
		IsMissing(co.Rs.PublicSubnet == nil, "public subnet"),
		IsMissing(co.Rs.PrivateSubnet == nil, "private subnet"),
		IsMissing(co.Rs.PublicRouteTable == nil, "main route table"),
		IsMissing(co.Rs.PrivateRouteTable == nil, "private route table"),
		IsMissing(co.Rs.InternetGateway == nil, "internet gateway"),
		IsMissing(co.Rs.SecurityGroups == nil, "security groups"),
		IsMissing(co.Rs.KeyPair == nil, "key pair"),
	}
	if util.AtLeastOneTrue(bools) {
		fmt.Println("run 'lab init' to create missing cloudlab resources")
		os.Exit(1)
	}
}

func IsMissing(missing bool, resourceType string) bool {
	if missing {
		fmt.Printf("%s is missing from cloudlab resources\n", resourceType)
	}
	return missing
}

// ######################################################################################
// Init #################################################################################
// ######################################################################################

// 1.  create key pair
// 2.  create vpc (auto creates public route table)
// 3.  create public subnet
// 4.  modify public subnet to enable IPv4 assignment
// 5.  create private subnet
// 6.  create private route table
// 7.  create internet gateway
// 8.  attach internet gateway to vpc
// 9.  add internet gateway to public route table
// 10. associate public route table with public subnet
// 11. create base security group for port 22 (tcp & udp)

func (co *AWSCloudOperator) InitializeCloudLabResources() {
	if co.Rs.KeyPair == nil {
		co.CreateKeyPair()
		co.Rs.KeyPair = co.Finder.findKeyPair()
	}
	if co.Rs.Vpc == nil {
		co.Rs.Vpc = co.Creator.createVpc(DefaultVpcCidrBlock, CloudLabVpc)
		co.Rs.PublicRouteTable = co.Finder.findMainRouteTable(co.Rs.Vpc)
	}
	if co.Rs.PublicSubnet == nil {
		co.Rs.PublicSubnet = co.Creator.createSubnet(co.Rs.Vpc, CloudLabPublicSubnet, PublicSubnetCidrBlock)
		co.resolvePublicSubnetAttributes()
	}
	if co.Rs.PrivateSubnet == nil {
		co.Rs.PrivateSubnet = co.Creator.createSubnet(co.Rs.Vpc, CloudLabPrivateSubnet, PrivateSubnetCidrBlock)
	}
	if co.Rs.PrivateRouteTable == nil {
		co.Rs.PrivateRouteTable = co.Creator.createRouteTable(co.Rs.Vpc, CloudLabPrivateRouteTable)
	}
	if co.Rs.InternetGateway == nil {
		co.Rs.InternetGateway = co.Creator.createInternetGateway(CloudLabInternetGateway)
	}
	if !internetGatewayIsAttachedToVpc(co.Rs.InternetGateway, co.Rs.Vpc) {
		attachInternetGatewayToVpc(co.Rs.InternetGateway, co.Rs.Vpc)
	}
	if !internetGatewayRouteExistsOnRouteTable(co.Rs.PublicRouteTable, co.Rs.InternetGateway) {
		addInternetGatewayRoute(co.Rs.PublicRouteTable, co.Rs.InternetGateway, RouteTablePublicSubnetCidr)
	}
	if !subnetAssociationExistsOnRouteTable(co.Rs.PublicRouteTable, co.Rs.PublicSubnet) {
		associateSubnetWithRouteTable(co.Rs.PublicRouteTable, co.Rs.PublicSubnet)
	}
	securityGroup22 := co.Finder.findSecurityGroupByName(co.Rs.SecurityGroups, "22")
	if securityGroup22 == nil {
		co.Creator.CreateSecurityGroup(co.Rs.Vpc, "22", 22)
		co.Rs.SecurityGroups = co.Finder.FindSecurityGroups()
	}
}

// ######################################################################################
// ######################################################################################
// ######################################################################################

func (co *AWSCloudOperator) DestroyCloudLabResources() {
	nonTerminatedInstances := co.Finder.FindNonTerminatedInstances()
	if len(nonTerminatedInstances) > 0 {
		panic("delete all instances before destroying cloudlab infrastructure")
	}

	if co.Rs.InternetGateway != nil {
		detachInternetGatewayFromVpc(co.Rs.InternetGateway, co.Rs.Vpc)
		co.Deleter.deleteInternetGateway(co.Rs.InternetGateway)
	}

	if co.Rs.PrivateRouteTable != nil {
		disassociateSubnetsFromRouteTable(co.Rs.PrivateRouteTable)
		co.Deleter.deleteRouteTable(co.Rs.PrivateRouteTable)
	}

	if co.Rs.PublicSubnet != nil {
		co.Deleter.deleteSubnet(co.Rs.PublicSubnet)
	}
	if co.Rs.PrivateSubnet != nil {
		co.Deleter.deleteSubnet(co.Rs.PrivateSubnet)
	}

	co.Deleter.deleteSecurityGroups(co.Rs.SecurityGroups)

	if co.Rs.Vpc != nil {
		co.Deleter.deleteVpc(co.Rs.Vpc)
	}

	if co.Rs.KeyPair != nil {
		co.Deleter.deleteKeyPair(co.Rs.KeyPair)
	}

	if util.ObjectExists(KeyFilePath()) {
		err := os.Remove(KeyFilePath())
		if err != nil {
			panic("failed to delete key file: " + KeyFilePath())
		}
	}
}
