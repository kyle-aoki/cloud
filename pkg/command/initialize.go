package command

import (
	"cloud/pkg/defaults"
)

func (c Commander) Initialize() {
	defaults.CreateCloudLabDefaults()
}
