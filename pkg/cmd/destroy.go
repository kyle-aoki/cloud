package cmd

import (
	"cloudlab/pkg/resource"
	"fmt"
)

func DestroyCloudLabResources() {
	ro := resource.NewResourceOperatorNoAudit()
	ro.DestroyCloudLabResources()
	fmt.Println("deleted all cloudlab resources")
}
