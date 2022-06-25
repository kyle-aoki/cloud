package cmd

import "cloudlab/pkg/resource"

func Info() {
	co := resource.NewCloudOperatorNoAudit()
	co.Info()
}
