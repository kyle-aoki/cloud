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

	co := resource.New()
	co.Rs.Instances = co.Finder.FindNonTerminatedInstances()

	var nameIds []NameId

	for _, inst := range co.Rs.Instances {
		instName := resource.FindNameTagValue(inst.Tags)
		util.Log("found instance '%s'", instName)
		if instName != nil && util.Contains(*instName, targets) {
			nameIds = append(nameIds, NameId{Name: *instName, Id: *inst.InstanceId})
		}
	}

	util.Log("nameIds %v", nameIds)

	for _, nameId := range nameIds {
		co.TerminateInstance(&nameId.Id)
		fmt.Println(nameId.Name)
	}
}
