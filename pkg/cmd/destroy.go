package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
)

func DestroyCloudLabResources() {
	util.Log("destroying cloudlab infrastructure...")
	co := resource.NewCloudOperatorNoAudit()
	co.DestroyCloudLabResources()
	fmt.Println("deleted all cloudlab infrastructure")
}
