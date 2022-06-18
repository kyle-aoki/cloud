package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"golang.org/x/sys/unix"
)

func ListKeys() {
	ro := resource.NewResourceOperator()
	for _, kp := range ro.KeyPairs {
		if kp.KeyName == nil {
			continue
		}
		fmt.Println(*kp.KeyName)
	}
}

const cloudlabConfigDirName = ".cloudlab"
const keyFile = "key.pem"

var keyFilePath string

func init() {
	homeDir, err := os.UserHomeDir()
	util.MustExec(err)
	switch runtime.GOOS {
	case "windows":
		keyFilePath = fmt.Sprintf("%s\\%s\\%s", homeDir, cloudlabConfigDirName, keyFile)
	default:
		keyFilePath = fmt.Sprintf("%s/%s/%s", homeDir, cloudlabConfigDirName, keyFile)
	}
}

func fatalInsufficientPermissions() {
	fmt.Println(fmt.Sprintf("cannot write key file to %s", cloudlabConfigDirName))
	fmt.Println("if on MacOS or Linux, try: 'sudo lab create key'")
	fmt.Println("if on windows, try running powershell or command prompt as an administrator'")
	os.Exit(1)
}

func writable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

func CreateKeyPair() {
	if !writable(cloudlabConfigDirName) {
		fatalInsufficientPermissions()
	}

	if !fsObjectExists(cloudlabConfigDirName) {
		createDir(cloudlabConfigDirName)
	}

	if !fsObjectExists(keyFilePath) {
		createFile(keyFilePath)
	}

	ro := resource.NewResourceOperator()
	keyMaterial := ro.ExecuteCreateKeyPairRequest()

	setFileContent(keyFilePath, *keyMaterial)

	fmt.Println(fmt.Sprintf("created key file at %s", keyFilePath))
	fmt.Println("to ssh with key file run:")
	fmt.Println()
	fmt.Println("lab create public node")
	fmt.Println()
	fmt.Println(fmt.Sprintf("ssh-add %s", keyFilePath))
	fmt.Println("ssh ubuntu@<public-ip-address>")
}

var test = 3 | 5

func setFileContent(keyFilePath, keyMaterial string) {
	f, err := os.OpenFile(keyFilePath, os.O_APPEND|os.O_WRONLY, 400)
	util.MustExec(err)
	defer f.Close()
	_, err = f.Write([]byte(keyMaterial))
	util.MustExec(err)
}

func createFile(keyFilePath string) {
	err := ioutil.WriteFile(keyFilePath, []byte{}, os.ModeAppend)
	util.MustExec(err)
}

func createDir(dir string) {
	err := os.Mkdir(dir, os.ModeDir)
	util.MustExec(err)
}

func fsObjectExists(fsObject string) bool {
	if _, err := os.Stat(fsObject); os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteKey() {
	resource.DeleteKey()
}

func DeleteAllKeys() {
}
