package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func ListInstances() {
	co := resource.New()

	if co.Finder.FindVpc(resource.CloudLabVpc) == nil {
		panic("run 'lab init' first")
	}

	if *args.Flags.All {
		PrintInstanceList(co.Finder.FindInstances())
	} else {
		PrintInstanceList(co.Finder.FindNonTerminatedInstances())
	}
}

func PrintInstanceList(instances []*ec2.Instance) {
	util.Tab("name\tstate\tprivate-ip\tpublic-ip\tports\tuptime")

	for _, inst := range instances {
		l := fmt.Sprintf("%v\t%v\t%v\t%v\t%s\t%s",
			Name(inst),
			State(inst),
			PrivateIp(inst),
			PublicIp(inst),
			Ports(inst),
			TimeElapsed(inst),
		)
		util.Tab(l)
	}

	util.ExecPrint()
}

func Ports(inst *ec2.Instance) (portsString string) {
	var ports []string
	for _, sg := range inst.SecurityGroups {
		if sg.GroupName == nil {
			panic("unknown security group found in instance")
		}
		ports = append(ports, *sg.GroupName)
	}
	sort.Strings(ports)
	return strings.Join(ports, ", ")
}

func Name(inst *ec2.Instance) string {
	n := resource.FindNameTagValue(inst.Tags)
	if n != nil {
		return *n
	}
	return ""
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

func TimeElapsed(inst *ec2.Instance) string {
	if inst.State.Name != nil && *inst.State.Name == "terminated" {
		return ""
	}
	d := time.Now().UnixNano() - inst.LaunchTime.UnixNano()
	f := time.Duration(d)
	switch {
	case f < time.Minute:
		return fmt.Sprintf("%vs", f.Round(time.Second).Seconds())
	case f < time.Hour:
		return fmt.Sprintf("%vm", f.Round(time.Minute).Minutes())
	case f < time.Hour*24:
		return fmt.Sprintf("%vh", f.Round(time.Hour).Hours())
	default:
		return fmt.Sprintf("%vd", int(f.Round(time.Hour).Hours())%24)
	}
}
