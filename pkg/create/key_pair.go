package create

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/config"
	"cloud/pkg/util"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func CreateKeyPair() {
	keyNames := args.Collect()

	for _, keyName := range keyNames {
		client := amazon.EC2Client()

		kpr := ec2.CreateKeyPairInput{
			KeyName: &keyName,
		}

		response, err := client.CreateKeyPair(&kpr)
		util.Check(err)

		keyFileName := fmt.Sprintf("%v/%v.pem", config.KeyDir(), keyName)

		err = ioutil.WriteFile(keyFileName, []byte(*response.KeyMaterial), 0400)
		util.Check(err)

		fmt.Println(keyFileName)
	}
}
