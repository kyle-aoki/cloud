package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"strings"
)

const InitGuideTemplate = `placed ssh key at <config-dir>/key.pem
create an instance: lab run
ssh into an instance: lab ssh <instance-name>
`

func InitializeCloudLabResources() {
	lr := resource.NewLabResources()
	createdKeyPair := lr.KeyPair == nil
	lr.CreateMissingResources()
	lr.Info()

	if createdKeyPair {
		initGuide := strings.ReplaceAll(InitGuideTemplate, "<config-dir>", util.ConfigDir())
		fmt.Println(initGuide)
	}
}
