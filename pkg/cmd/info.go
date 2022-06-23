package cmd

import "cloudlab/pkg/resource"

func Info() {
	ro := resource.NewResourceOperatorNoAudit()
	ro.Info()
}
