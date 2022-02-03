package config

import (
	"cloud/pkg/util"
)

func ReplaceVariables(ic InstanceConfiguration) InstanceConfiguration {
	ic.Subnet = replaceVariable(ic.Subnet, Vars.Variables.Subnets)
	for i, sg := range ic.SecurityGroups {
		ic.SecurityGroups[i] = replaceVariable(sg, Vars.Variables.SecurityGroups)
	}
	ic.Ami = replaceVariable(ic.Ami, Vars.Variables.AMIs)
	return ic
}

func replaceVariable(varName string, kvs []KeyValue) string {
	for _, kv := range kvs {
		if kv.Name == varName {
			return kv.Value
		}
	}
	util.PanicVerify("'%v' not found in variables.", varName)
	panic("")
}
