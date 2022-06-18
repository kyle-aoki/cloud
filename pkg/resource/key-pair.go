package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (ro *ResourceOperator) ExecuteCreateKeyPairRequest() *string {
	ckpo, err := amazon.EC2().CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName:           util.StrPtr(CloudLabKeyPair),
		TagSpecifications: CreateNameTagSpec("key-pair", CloudLabKeyPair),
	})
	util.MustExec(err)
	return ckpo.KeyMaterial
}


func (ro *ResourceOperator) CreateCloudlabKeyPair() {
	err := ioutil.WriteFile(util.ConfigDir(), []byte("test"), 400)
	util.MustExec(err)
	km := ro.ExecuteCreateKeyPairRequest()
	err = ioutil.WriteFile(util.ConfigDir(), []byte(*km), 400)
	util.MustExec(err)
}
