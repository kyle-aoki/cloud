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
	if resource.FindVpc() == nil {
		panic("run 'lab init' first")
	}

	if args.Flags.Quiet {
		printNamesOnly(resource.FindNonTerminatedInstances())
		return
	}

	if args.Flags.ShowTerminated {
		printInstanceList(resource.FindInstances())
	} else {
		printInstanceList(resource.FindNonTerminatedInstances())
	}
}

func printNamesOnly(instances []*ec2.Instance) {
	for _, inst := range instances {
		fmt.Printf("%v\n", name(inst))
	}
}

func printInstanceList(instances []*ec2.Instance) {
	util.TabPrint("name", "state", "private-ip", "public-ip", "ports", "uptime")
	for i := 0; i < len(instances); i++ {
		util.TabPrint(
			name(instances[i]),
			state(instances[i]),
			privateIp(instances[i]),
			publicIp(instances[i]),
			ports(instances[i]),
			timeElapsed(instances[i]),
		)
	}
	util.ExecPrint()
}

func ports(inst *ec2.Instance) (portsString string) {
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

func name(inst *ec2.Instance) string {
	n := resource.FindNameTagValue(inst.Tags)
	if n != nil {
		return *n
	}
	return ""
}

func state(inst *ec2.Instance) string {
	if inst.State.Name != nil {
		return *inst.State.Name
	}
	return ""
}

func publicIp(inst *ec2.Instance) string {
	if inst.PublicIpAddress != nil {
		return *inst.PublicIpAddress
	}
	return ""
}

func privateIp(inst *ec2.Instance) string {
	if inst.PrivateIpAddress != nil {
		return *inst.PrivateIpAddress
	}
	return ""
}

func timeElapsed(inst *ec2.Instance) string {
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
