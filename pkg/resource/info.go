package resource

import (
	"cloud/pkg/util"
	"fmt"
)

func (ro *ResourceOperator) Info() {
	if ro.Vpc == nil {
		util.Print("missing cloudlab vpc")
	} else {
		util.Print("vpc\t" + *ro.Vpc.VpcId)
	}

	if ro.PublicSubnet == nil {
		util.Print("missing cloudlab public subnet")
	} else {
		util.Print("public subnet\t" + *ro.PublicSubnet.SubnetId)
	}

	if ro.PrivateSubnet == nil {
		util.Print("missing cloudlab private subnet")
	} else {
		util.Print("private subnet\t" + *ro.PrivateSubnet.SubnetId)
	}

	if ro.PublicRouteTable == nil {
		util.Print("missing public route table")
	} else {
		util.Print("public route table\t" + *ro.PublicRouteTable.RouteTableId)
	}

	if ro.PrivateRouteTable == nil {
		util.Print("missing private route table")
	} else {
		util.Print("private route table\t" + *ro.PrivateRouteTable.RouteTableId)
	}

	if ro.InternetGateway == nil {
		util.Print("missing cloudlab internet gateway")
	} else {
		util.Print("internet gateway\t" + *ro.InternetGateway.InternetGatewayId)
	}

	util.Print("\t")
	for _, sg := range ro.SecurityGroups {
		util.Print("security group " + *sg.GroupName + "\t" + *sg.GroupId)
	}
	util.Print("\t")

	util.Print("number of key pairs\t" + util.IntToString(len(ro.KeyPairs)))
	if ro.CurrentKeyPair != nil {
		util.Print("current key pair\t" + *ro.CurrentKeyPair.KeyName)
	}

	util.Print("\t")
	util.Print("number of nodes\t" + util.IntToString(len(ro.Instances)))

	util.Print("\t")
	for _, inst := range ro.Instances {
		util.Print(fmt.Sprintf(
			"%s\t%s\t%s",
			*inst.InstanceId,
			*FindNameTagValue(inst.Tags),
			*inst.State.Name,
		))
	}
	util.Print("\t")

	util.Flush()
}
