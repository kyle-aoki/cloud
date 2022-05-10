package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/defaults"
	"cloud/pkg/util"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetIds() []string {
	cldo := defaults.Start()
	names := args.Collect()

	nodeNames := GetNodeNames(cldo.Instances)
	var nodesToStart []NodeName

	for i := range names {
		for j := range nodeNames {
			if strings.Contains(nodeNames[j].Name, names[i]) {
				nodesToStart = append(nodesToStart, nodeNames[j])
			}
		}
	}

	var ids []string
	for i := range nodesToStart {
		ids = append(ids, nodesToStart[i].Id)
	}

	return ids
}

func Stop() {
	ids := GetIds()
	client := amazon.EC2()

	sii, err := client.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: util.StrSlicePtr(ids),
	})
	util.MustExec(err)

	for i := range sii.StoppingInstances {
		fmt.Println(*sii.StoppingInstances[i].InstanceId)
	}
}

func Start() {
	ids := GetIds()
	client := amazon.EC2()

	sio, err := client.StartInstances(&ec2.StartInstancesInput{
		InstanceIds: util.StrSlicePtr(ids),
	})
	util.MustExec(err)

	for i := range sio.StartingInstances {
		fmt.Println(*sio.StartingInstances[i].InstanceId)
	}
}

type NodeName struct {
	Name string
	Id   string
}

func GetNodeNames(nodes []*ec2.Instance) (nodeNames []NodeName) {
	for i := 0; i < len(nodes); i++ {
		var nameTag string
		for j := range nodes[i].Tags {
			if *nodes[i].Tags[j].Key == "Name" {
				nameTag = *nodes[i].Tags[j].Value
			}
		}
		nodeNames = append(nodeNames, NodeName{
			Id:   *nodes[i].InstanceId,
			Name: nameTag,
		})
	}
	return nodeNames
}
