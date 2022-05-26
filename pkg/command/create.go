package command

import "cloud/pkg/resource"

func CreatePublicInstance() {
	_ = resource.NewResourceOperator()

}

func CreatePrivateInstance() {
	_ = resource.NewResourceOperator()

}
