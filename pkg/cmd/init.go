package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"strings"
)

const InitGuideTemplate = `
placed ssh key at <config-dir>

create an instance:

lab run

ssh into an instance:

ssh -i <config-dir> ubuntu@<public-ip>`

func InitializeCloudLabResources() {
	co := resource.NewCloudOperatorNoAudit()
	co.InitializeCloudLabResources()
	co.Info()

	initGuide := strings.ReplaceAll(InitGuideTemplate, "<config-dir", util.ConfigDir())

	fmt.Println(initGuide)
}
