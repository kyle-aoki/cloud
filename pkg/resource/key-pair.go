package resource

import (
	"cloudlab/pkg/util"
	"fmt"
	"log"
	"os"

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

// ##############################################################################
// Main Logic ###################################################################
// ##############################################################################

func (co *AWSCloudOperator) CreateKeyPair() {
	log.Print("starting key pair creation process")

	if util.ObjectExists(keyFilePath()) {
		util.DeleteFile(keyFilePath())
	}
	if !util.ObjectExists(util.ConfigDir()) {
		util.CreateDir(util.ConfigDir())
	}
	if !writable(util.ConfigDir()) {
		fatalInsufficientPermissions()
	}
	if !util.ObjectExists(util.ConfigDir()) {
		util.CreateDir(util.ConfigDir())
	}
	if !util.ObjectExists(keyFilePath()) {
		util.CreateEmptyFile(keyFilePath())
	}

	keyMaterial := co.creator.ExecuteCreateKeyPairRequest(CloudLabKeyPair)

	setFileContentWith400(keyFilePath(), *keyMaterial)
}

// ##############################################################################
// Key Writing Utils ############################################################
// ##############################################################################

func setFileContentWith400(path, text string) {
	log.Println("setting file content")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 400)
	util.MustExec(err)
	defer f.Close()
	_, err = f.Write([]byte(text))
	util.MustExec(err)
}

func writable(path string) bool {
	isWritable := unix.Access(path, unix.W_OK) == nil
	log.Println(path, "isWritable", isWritable)
	return isWritable
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
