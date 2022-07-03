package resource

import (
	"cloudlab/pkg/util"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

const keyFileName = "key.pem"

func KeyFilePath() string {
	return fmt.Sprintf("%s/%s", util.ConfigDir(), keyFileName)
}

func (co *AWSCloudOperator) CreateKeyPair() {
	log.Println("starting key pair creation process")

	log.Println("attempting to remove everything in config dir")
	err := os.RemoveAll(util.ConfigDir())
	util.MustExec(err)

	log.Println("checking if home dir is writable: ", util.HomeDir())
	if !writable(util.HomeDir()) {
		fatalInsufficientPermissions()
	}

	log.Println("creating config dir: ", util.ConfigDir())
	util.CreateDir(util.ConfigDir())

	keyMaterial := co.Creator.ExecuteCreateKeyPairRequest(CloudLabKeyPair)

	log.Println("writing key material to key file at", KeyFilePath())
	err = ioutil.WriteFile(KeyFilePath(), []byte(*keyMaterial), 0400)
	util.MustExec(err)

	log.Println("changing key file permissions to 400")
	err = os.Chmod(KeyFilePath(), 0400)
	util.MustExec(err)
}

func writable(path string) bool {
	isWritable := unix.Access(path, unix.W_OK) == nil
	log.Println(path, "isWritable", isWritable)
	return isWritable
}

func fatalInsufficientPermissions() {
	fmt.Println(fmt.Sprintf("cannot write key file to %s", KeyFilePath()))
	fmt.Println("if on MacOS or Linux, try: 'sudo lab create key'")
	fmt.Println("if on Windows, try running powershell or command prompt as an administrator'")
	os.Exit(1)
}
