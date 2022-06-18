package cmd

import "cloudlab/pkg/resource"

func Info() {
	ro := resource.New()
	ro.FindAll()
	ro.Info()
}
