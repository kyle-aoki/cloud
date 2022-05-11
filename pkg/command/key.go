package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/defaults"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
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

func DeleteKey() {
	keyPairNames := args.Collect()
	DeleteKeys(keyPairNames)
}

func DeleteAllKeys() {
	cldo := defaults.NewOperator()
	cldo.FindAllCloudLabKeyPairs()
	var keyPairNames []string
	for _, kp := range cldo.KeyPairs {
		keyPairNames = append(keyPairNames, *kp.KeyName)
	}
	DeleteKeys(keyPairNames)
}

func DeleteKeys(keyPairNames []string) {
	for _, keyPairName := range keyPairNames {
		_, err := amazon.EC2().DeleteKeyPair(&ec2.DeleteKeyPairInput{
			KeyName: util.StrPtr(keyPairName),
		})
		util.MustExec(err)
		util.VMessage("deleted", defaults.CloudLabKeyPair, keyPairName)
	}
}
