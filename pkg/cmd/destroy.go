package cmd

import (
	"cloudlab/pkg/resource"
	"fmt"
)

func DestroyCloudLabResources() {
	co := resource.NewCloudOperatorNoAudit()
	co.DestroyCloudLabResources()
	fmt.Println("deleted all cloudlab resources")
}
