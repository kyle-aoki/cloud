package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"strings"
)

const InitGuideTemplate = `placed ssh key at <config-dir>/key.pem
create an instance: lab run
ssh into an instance: ssh -i <config-dir>/key.pem ubuntu@<public-ip>
`

func InitializeCloudLabResources() {
	co := resource.NewCloudOperatorNoAudit()
	createdSSH := co.Rs.KeyPair == nil
	co.InitializeCloudLabResources()
	co.Info()

	if createdSSH {
		initGuide := strings.ReplaceAll(InitGuideTemplate, "<config-dir>", util.ConfigDir())
		fmt.Println(initGuide)
	}
}
