package resource

import (
	"cloudlab/pkg/util"
)

func (co *AWSCloudOperator) Info() {
	if co.Rs.Vpc == nil {
		util.Tab("missing cloudlab vpc")
	} else {
		util.Tab("vpc\t" + *co.Rs.Vpc.VpcId)
	}

	if co.Rs.PublicSubnet == nil {
		util.Tab("missing cloudlab public subnet")
	} else {
		util.Tab("public subnet\t" + *co.Rs.PublicSubnet.SubnetId)
	}

	if co.Rs.PrivateSubnet == nil {
		util.Tab("missing cloudlab private subnet")
	} else {
		util.Tab("private subnet\t" + *co.Rs.PrivateSubnet.SubnetId)
	}

	if co.Rs.PublicRouteTable == nil {
		util.Tab("missing public route table")
	} else {
		util.Tab("public route table\t" + *co.Rs.PublicRouteTable.RouteTableId)
	}

	if co.Rs.PrivateRouteTable == nil {
		util.Tab("missing private route table")
	} else {
		util.Tab("private route table\t" + *co.Rs.PrivateRouteTable.RouteTableId)
	}

	if co.Rs.InternetGateway == nil {
		util.Tab("missing cloudlab internet gateway")
	} else {
		util.Tab("internet gateway\t" + *co.Rs.InternetGateway.InternetGatewayId)
	}

	if co.Rs.KeyPair == nil {
		util.Tab("missing cloudlab key")
	} else {
		util.Tab("key\t" + *co.Rs.KeyPair.KeyPairId)
	}

	for _, sg := range co.Rs.SecurityGroups {
		util.Tab("security group " + *sg.GroupName + "\t" + *sg.GroupId)
	}

	util.ExecPrint()
}
