package defaults

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// const KeyPairHelpText = `you are creating a cloudlab keypair
// pay close attention
// the key material will be printed out to the console
// copy and paste the key material into a file called 'cloudlab-key.pem'
// place this file inside the Documents folder
// placing this file in an insecure folder may cause problems
// for MacOS & Linux users, run the following command on your key file:
// > chmod 400 ~/Documents/cloudlab-key.pem
// this will set the correct file permissions
// on Windows, placing the key in the Documents folder should be enough
// to use the key, run the following command:
// > ssh -i ~/Documents/keys/cloudlab-key.pem ubuntu@<public-ip-address>
// alternatively, you can save the key with 'ssh-add', and connect without the key:
// > ssh-add ~/Documents/keys/cloudlab-key.pem
// > ssh ubuntu@<public-ip-address>
// you should only use one cloudlab keypair at a time
// creating an instance will use the latest key-pair created
// to create a key-pair without seeing this text, use the -q argument
// ok
// here is the key material:`

func InitiateKeyPairCreation() {
	args.ParseKeyPairFlags()
	cldo := &CloudLabDefaultsOperator{}
	cldo.FindAllCloudLabKeyPairs()
	fmt.Println(*cldo.createKeyPair())
}

func (cldo *CloudLabDefaultsOperator) FindAllCloudLabKeyPairs() {
	dkpo, err := amazon.EC2().DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	util.MustExec(err)

	for _, kp := range dkpo.KeyPairs {
		if nameTagEquals(kp.Tags, CloudLabKeyPair) {
			cldo.KeyPairs = append(cldo.KeyPairs, kp)
		}
	}

	for _, kp := range cldo.KeyPairs {
		if kp.KeyName != nil && *kp.KeyName == cldo.GetCurrentCloudLabKeyPairName() {
			cldo.CurrentKeyPair = kp
			break
		}
	}
}

func (cldo *CloudLabDefaultsOperator) GetCurrentCloudLabKeyPairName() string {
	number := len(cldo.KeyPairs)
	return fmt.Sprintf("%s%v", CloudLabKeyPairNameTemplate, number)
}

func (cldo *CloudLabDefaultsOperator) getNextCloudLabKeyPairName() string {
	number := len(cldo.KeyPairs) + 1
	return fmt.Sprintf("%s%v", CloudLabKeyPairNameTemplate, number)
}

func (cldo *CloudLabDefaultsOperator) createKeyPair() *string {
	ckpo, err := amazon.EC2().CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName:           util.StrPtr(cldo.getNextCloudLabKeyPairName()),
		TagSpecifications: CreateNameTagSpec("key-pair", CloudLabKeyPair),
	})
	util.MustExec(err)
	return ckpo.KeyMaterial
}
