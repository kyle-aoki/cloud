package command

import "cloud/pkg/initialize"

func Initialize() {
	initialize.CreateCloudLabDefaults()
}
