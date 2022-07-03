package resource

import (
	"cloudlab/pkg/util"
	"fmt"
)

func (co *AWSCloudOperator) Info() {
	if co.Rs.Vpc == nil {
		util.SetTabPrint("missing cloudlab vpc")
	} else {
		util.SetTabPrint("vpc\t" + *co.Rs.Vpc.VpcId)
	}

	if co.Rs.PublicSubnet == nil {
		util.SetTabPrint("missing cloudlab public subnet")
	} else {
		util.SetTabPrint("public subnet\t" + *co.Rs.PublicSubnet.SubnetId)
	}

	if co.Rs.PrivateSubnet == nil {
		util.SetTabPrint("missing cloudlab private subnet")
	} else {
		util.SetTabPrint("private subnet\t" + *co.Rs.PrivateSubnet.SubnetId)
	}

	if co.Rs.PublicRouteTable == nil {
		util.SetTabPrint("missing public route table")
	} else {
		util.SetTabPrint("public route table\t" + *co.Rs.PublicRouteTable.RouteTableId)
	}

	if co.Rs.PrivateRouteTable == nil {
		util.SetTabPrint("missing private route table")
	} else {
		util.SetTabPrint("private route table\t" + *co.Rs.PrivateRouteTable.RouteTableId)
	}

	if co.Rs.InternetGateway == nil {
		util.SetTabPrint("missing cloudlab internet gateway")
	} else {
		util.SetTabPrint("internet gateway\t" + *co.Rs.InternetGateway.InternetGatewayId)
	}

	if co.Rs.KeyPair == nil {
		util.SetTabPrint("missing cloudlab key")
	} else {
		util.SetTabPrint("key\t" + *co.Rs.KeyPair.KeyPairId)
	}

	util.SetTabPrint("\t")
	for _, sg := range co.Rs.SecurityGroups {
		util.SetTabPrint("security group " + *sg.GroupName + "\t" + *sg.GroupId)
	}

	if len(co.Rs.SecurityGroups) > 0 {
		util.SetTabPrint("\t")
	}

	for _, inst := range co.Rs.Instances {
		util.SetTabPrint(fmt.Sprintf(
			"%s\t%s",
			*FindNameTagValue(inst.Tags),
			*inst.State.Name,
		))
	}
	util.SetTabPrint("\t")

	util.TabPrint()
}
