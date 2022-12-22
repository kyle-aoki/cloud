package resource

import (
	"cloudlab/pkg/util"
)

func PrintInfo(lr *LabResources) {
	if lr.Vpc == nil {
		util.Tab("missing cloudlab vpc")
	} else {
		util.Tab("vpc\t" + *lr.Vpc.VpcId)
	}

	if lr.PublicSubnet == nil {
		util.Tab("missing cloudlab public subnet")
	} else {
		util.Tab("public subnet\t" + *lr.PublicSubnet.SubnetId)
	}

	if lr.PrivateSubnet == nil {
		util.Tab("missing cloudlab private subnet")
	} else {
		util.Tab("private subnet\t" + *lr.PrivateSubnet.SubnetId)
	}

	if lr.PublicRouteTable == nil {
		util.Tab("missing public route table")
	} else {
		util.Tab("public route table\t" + *lr.PublicRouteTable.RouteTableId)
	}

	if lr.PrivateRouteTable == nil {
		util.Tab("missing private route table")
	} else {
		util.Tab("private route table\t" + *lr.PrivateRouteTable.RouteTableId)
	}

	if lr.InternetGateway == nil {
		util.Tab("missing cloudlab internet gateway")
	} else {
		util.Tab("internet gateway\t" + *lr.InternetGateway.InternetGatewayId)
	}

	if lr.KeyPair == nil {
		util.Tab("missing cloudlab key")
	} else {
		util.Tab("key\t" + *lr.KeyPair.KeyPairId)
	}

	for _, sg := range lr.SecurityGroups {
		util.Tab("security group " + *sg.GroupName + "\t" + *sg.GroupId)
	}

	util.ExecPrint()
}
