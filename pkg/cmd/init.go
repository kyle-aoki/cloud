package cmd

import "cloudlab/pkg/resource"

func InitializeCloudLabResources() {
	ro := resource.New()
	ro.FindAll()
	ro.InitializeCloudLabResources()
	ro.Info()
}
