package command

import "cloud/pkg/resource"

func InitializeCloudLabResources() {
	ro := resource.New()
	ro.FindAll()
	ro.InitializeCloudLabResources()
	ro.Info()
}
