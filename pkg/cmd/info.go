package cmd

import "cloudlab/pkg/resource"

func Info() {
	lr := resource.FindAllLabResources()
	lr.Info()
}
