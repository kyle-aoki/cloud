package command

import (
	"cloud/pkg/defaults"
	"cloud/pkg/util"
	"fmt"
)

func Info() {
	cldo := defaults.Start()

	if cldo.Vpc == nil {
		fmt.Println("missing cloudlab vpc")
	} else {
		util.Print("vpc\t" + *cldo.Vpc.VpcId)
	}

	if cldo.PublicSubnet == nil {
		fmt.Println("missing cloudlab public subnet")
	} else {
		util.Print("public subnet\t" + *cldo.PublicSubnet.SubnetId)
	}

	if cldo.PrivateSubnet == nil {
		fmt.Println("missing cloudlab private subnet")
	} else {
		util.Print("private subnet\t" + *cldo.PrivateSubnet.SubnetId)
	}

	if cldo.InternetGateway == nil {
		fmt.Println("missing cloudlab internet gateway")
	} else {
		util.Print("internet gateway\t" + *cldo.InternetGateway.InternetGatewayId)
	}

	if cldo.RouteTable == nil {
		fmt.Println("missing cloudlab route table")
	} else {
		util.Print("route table\t" + *cldo.RouteTable.RouteTableId)
	}

	util.Print("\t")
	for _, sg := range cldo.SecurityGroups {
		util.Print("security group " + *sg.GroupName + "\t" + *sg.GroupId)
	}
	util.Print("\t")

	util.Print(fmt.Sprintf("number of key pairs\t%d", len(cldo.KeyPairs)))
	if cldo.CurrentKeyPair == nil {
		fmt.Println("no cloudlab keypairs found")
	} else {
		util.Print("current key pair\t" + *cldo.CurrentKeyPair.KeyName)
	}

	util.Print("\t")
	util.Print(fmt.Sprintf("number of nodes\t%d", len(cldo.Instances)))

	util.Print("\t")
	for _, inst := range cldo.Instances {
		util.Print(fmt.Sprintf(
			"%s\t%s\t%s",
			*inst.InstanceId,
			*defaults.FindNameTagValue(inst.Tags),
			*inst.State.Name,
		))
	}
	util.Print("\t")

	util.Flush()
}
