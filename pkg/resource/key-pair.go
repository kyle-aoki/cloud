package resource

import (
	"cloudlab/pkg/util"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

const keyFileName = "key.pem"

func KeyFilePath() string {
	return fmt.Sprintf("%s/%s", util.ConfigDir(), keyFileName)
}

func ExecuteKeyPairCreationProcess() {
	log.Println("starting key pair creation process")

	log.Println("attempting to remove everything in config dir")
	err := os.RemoveAll(util.ConfigDir())
	util.Check(err)

	log.Println("checking if home dir is writable: ", util.HomeDir())
	if !isWritable(util.HomeDir()) {
		fatalInsufficientPermissions()
	}

	log.Println("creating config dir: ", util.ConfigDir())
	util.CreateDir(util.ConfigDir())

	keyMaterial := createKeyPair(CloudLabKeyPair)

	log.Println("writing key material to key file at", KeyFilePath())
	err = os.WriteFile(KeyFilePath(), []byte(*keyMaterial), 0400)
	util.Check(err)

	log.Println("changing key file permissions to 400")
	err = os.Chmod(KeyFilePath(), 0400)
	util.Check(err)
}

func isWritable(path string) (b bool) {
	b = unix.Access(path, unix.W_OK) == nil
	log.Println(path, "isWritable", b)
	return b
}

func fatalInsufficientPermissions() {
	fmt.Printf("cannot write key file to %s", KeyFilePath())
	fmt.Println("if on MacOS or Linux, try: 'sudo lab init'")
	fmt.Println("if on Windows, try running powershell or command prompt as an administrator'")
	os.Exit(1)
}
