package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/util"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetCurrentCloudLabKeyPairName(kps []*ec2.KeyPairInfo) string {
	number := len(kps)
	return fmt.Sprintf("%s%v", CloudLabKeyPairNameTemplate, number)
}

func (ro *ResourceOperator) getNextCloudLabKeyPairName() string {
	number := len(ro.KeyPairs) + 1
	return fmt.Sprintf("%s%v", CloudLabKeyPairNameTemplate, number)
}

func (ro *ResourceOperator) ExecuteCreateKeyPairRequest() *string {
	ckpo, err := amazon.EC2().CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName:           util.StrPtr(CloudLabKeyPair),
		TagSpecifications: CreateNameTagSpec("key-pair", CloudLabKeyPair),
	})
	util.MustExec(err)
	return ckpo.KeyMaterial
}

func DeleteKey() {
	keyPairNames := args.Collect()
	DeleteKeys(keyPairNames)
}

func (ro *ResourceOperator) GetKeyPairNames() (keyPairNames []string) {
	for _, kp := range ro.KeyPairs {
		keyPairNames = append(keyPairNames, *kp.KeyName)
	}
	return keyPairNames
}

// func DeleteAllKeys() {
// 	ro := NewOperator()
// 	FindAllCloudLabKeyPairs()
// 	DeleteKeys(ro.GetKeyPairNames())
// }

func DeleteKeys(keyPairNames []string) {
	for _, keyPairName := range keyPairNames {
		_, err := amazon.EC2().DeleteKeyPair(&ec2.DeleteKeyPairInput{
			KeyName: util.StrPtr(keyPairName),
		})
		util.MustExec(err)
		util.VMessage("deleted", CloudLabKeyPair, keyPairName)
	}
}

func (ro *ResourceOperator) CreateCloudlabKeyPair() {
	err := ioutil.WriteFile(util.ConfigDir(), []byte("test"), 400)
	util.MustExec(err)
	km := ro.ExecuteCreateKeyPairRequest()
	err = ioutil.WriteFile(util.ConfigDir(), []byte(*km), 400)
	util.MustExec(err)
}
