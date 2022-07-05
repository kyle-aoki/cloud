package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
)

type NameId struct {
	Name string
	Id   string
}

func DeleteInstances() {
	targets := args.Collect()
	util.Log("found delete targets: %v", targets)

	lr := resource.NewLabResources()
	lr.Instances = resource.FindNonTerminatedInstances()

	var nameIds []NameId

	for _, inst := range lr.Instances {
		instName := resource.FindNameTagValue(inst.Tags)
		util.Log("found instance '%s'", instName)
		if instName != nil && util.Contains(*instName, targets) {
			nameIds = append(nameIds, NameId{Name: *instName, Id: *inst.InstanceId})
		}
	}

	util.Log("nameIds %v", nameIds)

	for _, nameId := range nameIds {
		resource.TerminateInstance(&nameId.Id)
		fmt.Println(nameId.Name)
	}
}
