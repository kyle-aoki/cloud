package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
)

func DestroyCloudLabResources() {
	util.Log("destroying cloudlab infrastructure...")
	lr := resource.FindAllLabResources()
	resource.DestroyCloudLabResources(lr)
	fmt.Println("deleted all cloudlab infrastructure")
}
