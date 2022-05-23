package command

import (
	"cloud/pkg/defaults"
	"fmt"
)

func ListKeys() {
	cldo := defaults.NewOperator()
	cldo.FindAllCloudLabKeyPairs()
	for _, kp := range cldo.KeyPairs {
		if kp.KeyName == nil {
			continue
		}
		fmt.Println(*kp.KeyName)
	}
}

func CreateKeyPair() {
	cldo := defaults.Start()
	cldo.FindAllCloudLabKeyPairs()
	keyMaterial := cldo.CreateKeyPair()
	fmt.Println(*keyMaterial)
}

func DeleteKey() {
	defaults.DeleteKey()
}

func DeleteAllKeys() {
	defaults.DeleteAllKeys()
}
