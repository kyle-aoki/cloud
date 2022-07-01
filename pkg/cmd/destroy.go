package cmd

import (
	"cloudlab/pkg/resource"
	"fmt"
)

func DestroyCloudLabResources() {
	co := resource.NewCloudOperatorNoAudit()
	if len(co.Rs.Instances) != 0 {
		panic("cloudlab instances still exist\nplease delete them before running 'lab destroy'")
	}
	co.DestroyCloudLabResources()
	fmt.Println("deleted all cloudlab resources")
}
