package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/config"
	"cloud/pkg/tab"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetNodes() (nodes []*ec2.Instance) {
	client := amazon.EC2Client()

	di, err := client.DescribeInstances(&ec2.DescribeInstancesInput{})
	util.MustExec(err)

	for _, res := range di.Reservations {
		for _, inst := range res.Instances {
			nodes = append(nodes, inst)
		}
	}

	return nodes
}

func PrintNodes() {
	nodes := GetNodes()

	tab.Print("id\tname\tstate\tprivate-ip\tpublic-ip")

	for _, node := range nodes {
		if !config.Vars.ShowTerminatedNodes && State(node) == "terminated" {
			continue
		}
		l := fmt.Sprintf("%v\t%v\t%v\t%v\t%v",
			Id(node), Name(node), State(node), PrivateIp(node), PublicIp(node),
		)
		tab.Print(l)
	}
	tab.Flush()
}

func State(inst *ec2.Instance) string {
	if inst.State.Name != nil {
		return *inst.State.Name
	}
	return ""
}

func PublicIp(inst *ec2.Instance) string {
	if inst.PublicIpAddress != nil {
		return *inst.PublicIpAddress
	}
	return ""
}

func PrivateIp(inst *ec2.Instance) string {
	if inst.PrivateIpAddress != nil {
		return *inst.PrivateIpAddress
	}
	return ""
}

func Name(inst *ec2.Instance) string {
	for _, tag := range inst.Tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}
	return ""
}

func Id(inst *ec2.Instance) string {
	return *inst.InstanceId
}
