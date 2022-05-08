package command

import "cloud/pkg/defaults"

func Destroy() {
	defaults.DestroyCloudLabResources()
}
