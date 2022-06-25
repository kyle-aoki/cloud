package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
	"golang.org/x/sys/unix"
)

// ##############################################################################
// Key File #####################################################################
// ##############################################################################

const keyFileName = "key.pem"

func keyFilePath() string {
	return fmt.Sprintf("%s/%s", util.ConfigDir(), keyFileName)
}

// ##############################################################################
// AWS ##########################################################################
// ##############################################################################

func (co *AWSCloudOperator) ExecuteCreateKeyPairRequest() *string {
	ckpo, err := amazon.EC2().CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName:           util.StrPtr(CloudLabKeyPair),
		TagSpecifications: CreateNameTagSpec("key-pair", CloudLabKeyPair),
	})
	util.MustExec(err)
	return ckpo.KeyMaterial
}

// ##############################################################################
// Main Logic ###################################################################
// ##############################################################################

func CreateKeyPair() {
	if !writable(util.ConfigDir()) {
		fatalInsufficientPermissions()
	}
	if !util.ObjectExists(util.ConfigDir()) {
		util.CreateDir(util.ConfigDir())
	}
	if !util.ObjectExists(keyFilePath()) {
		util.CreateEmptyFile(keyFilePath())
	}

	co := NewCloudOperator()
	keyMaterial := co.ExecuteCreateKeyPairRequest()

	setFileContentWith400(keyFilePath(), *keyMaterial)

	fmt.Println(fmt.Sprintf(`
	created key file at %s
	
	to ssh with key file run:
	
	lab create public node
	
	ssh-add %s
	
	ssh ubuntu@<public-ip-address>`, keyFilePath(), keyFilePath()))
}

// ##############################################################################
// Key Writing Utils ############################################################
// ##############################################################################

func setFileContentWith400(path, text string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 400)
	util.MustExec(err)
	defer f.Close()
	_, err = f.Write([]byte(text))
	util.MustExec(err)
}

func writable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

// ##############################################################################
// Error Message ################################################################
// ##############################################################################

func fatalInsufficientPermissions() {
	fmt.Println(fmt.Sprintf("cannot write key file to %s", keyFilePath()))
	fmt.Println("if on MacOS or Linux, try: 'sudo lab create key'")
	fmt.Println("if on Windows, try running powershell or command prompt as an administrator'")
	os.Exit(1)
}
