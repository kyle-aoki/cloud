package command

import (
	"cloud/pkg/defaults"
	"cloud/pkg/tab"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func ListInstances() {
	cldo := defaults.Start()

	tab.Print("name\tstate\tprivate-ip\tpublic-ip\tports")

	for _, node := range cldo.Instances {
		if *node.State.Name == "terminated" {
			continue
		}
		l := fmt.Sprintf("%v\t%v\t%v\t%v\t%s",
			Name(node),
			State(node),
			PrivateIp(node),
			PublicIp(node),
			Ports(node),
		)
		tab.Print(l)
	}
	tab.Flush()
}

func Ports(inst *ec2.Instance) (ports string) {
	for i, sg := range inst.SecurityGroups {
		if sg.GroupName == nil {
			panic("unknown security group found in instance")
		}
		if i == len(inst.SecurityGroups)-1 {
			ports += *sg.GroupName
			break
		}
		ports += *sg.GroupName + ", "
	}
	return ports
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
