package command

import "cloud/pkg/defaults"

func (c Commander) Destroy() {
	defaults.DestroyCloudLabResources()
}
