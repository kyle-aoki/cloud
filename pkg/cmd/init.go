package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"strings"
)

const InitGuideTemplate = `placed ssh key at <config-dir>/key.pem
create an instance: lab run
ssh into an instance: ssh <instance-name>
`

func InitializeCloudLabResources() {
	lr := resource.NewLabResources()
	createdKeyPair := lr.KeyPair == nil
	resource.CreateMissingResources(lr)
	resource.PrintInfo(lr)

	if createdKeyPair {
		initGuide := strings.ReplaceAll(InitGuideTemplate, "<config-dir>", util.ConfigDir())
		fmt.Println(initGuide)
	}
}
