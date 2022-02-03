package create

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/util"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func CreateKeyPair() {
	keyName := args.Poll()
	
	client := amazon.EC2Client()

	kpr := ec2.CreateKeyPairInput{
		KeyName: &keyName,
	}

	response, err := client.CreateKeyPair(&kpr)
	util.Check(err)

	keyFileName := fmt.Sprintf("%v.pem", keyName)

	ioutil.WriteFile(keyFileName, []byte(*response.KeyMaterial), 0400)
	fmt.Println("created " + keyFileName)
}
