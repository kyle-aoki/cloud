package get

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetNodes() (nodes []*ec2.Instance) {
	client := amazon.EC2Client()

	di, err := client.DescribeInstances(&ec2.DescribeInstancesInput{})
	util.Check(err)

	for _, res := range di.Reservations {
		for _, inst := range res.Instances {
			nodes = append(nodes, inst)
		}
	}

	return nodes
}

func PrintNodes() {
	nodes := GetNodes()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	fmt.Fprintln(w, "name\tstate\tprivate-ip\tpublic-ip")

	for _, node := range nodes {
		l := fmt.Sprintf("%v\t%v\t%v\t%v",
			Name(node), State(node), PrivateIp(node), PublicIp(node),
		)
		fmt.Fprintln(w, l)
	}

	w.Flush()
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
