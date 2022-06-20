package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
)

func InitializeCloudLabResources() {
	ro := resource.New()
	ro.FindAll()
	ro.InitializeCloudLabResources()
	ro.Info()
	fmt.Println(fmt.Sprintf("placed ssh key at %s", util.ConfigDir()))
	fmt.Println()
	fmt.Println("create an instance:")
	fmt.Println()
	fmt.Println("lab run")
	fmt.Println()
	fmt.Println("ssh into an instance:")
	fmt.Println()
	fmt.Println(fmt.Sprintf("ssh -i %s ubuntu@<public-ip>", util.ConfigDir()))
	fmt.Println()
	fmt.Println()
}
