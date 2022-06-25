package resource

import (
	"cloudlab/pkg/util"
	"fmt"
)

func (co *AWSCloudOperator) Info() {
	if co.Rs.Vpc == nil {
		util.Print("missing cloudlab vpc")
	} else {
		util.Print("vpc\t" + *co.Rs.Vpc.VpcId)
	}

	if co.Rs.PublicSubnet == nil {
		util.Print("missing cloudlab public subnet")
	} else {
		util.Print("public subnet\t" + *co.Rs.PublicSubnet.SubnetId)
	}

	if co.Rs.PrivateSubnet == nil {
		util.Print("missing cloudlab private subnet")
	} else {
		util.Print("private subnet\t" + *co.Rs.PrivateSubnet.SubnetId)
	}

	if co.Rs.PublicRouteTable == nil {
		util.Print("missing public route table")
	} else {
		util.Print("public route table\t" + *co.Rs.PublicRouteTable.RouteTableId)
	}

	if co.Rs.PrivateRouteTable == nil {
		util.Print("missing private route table")
	} else {
		util.Print("private route table\t" + *co.Rs.PrivateRouteTable.RouteTableId)
	}

	if co.Rs.InternetGateway == nil {
		util.Print("missing cloudlab internet gateway")
	} else {
		util.Print("internet gateway\t" + *co.Rs.InternetGateway.InternetGatewayId)
	}

	if co.Rs.KeyPair == nil {
		util.Print("missing cloudlab key")
	} else {
		util.Print("key\t" + *co.Rs.KeyPair.KeyPairId)
	}

	util.Print("\t")
	for _, sg := range co.Rs.SecurityGroups {
		util.Print("security group " + *sg.GroupName + "\t" + *sg.GroupId)
	}

	if len(co.Rs.SecurityGroups) > 0 {
		util.Print("\t")
	}

	for _, inst := range co.Rs.Instances {
		util.Print(fmt.Sprintf(
			"%s\t%s",
			*FindNameTagValue(inst.Tags),
			*inst.State.Name,
		))
	}
	util.Print("\t")

	util.Flush()
}
