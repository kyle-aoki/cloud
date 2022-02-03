package run

import (
	"cloud/pkg/config"
	"cloud/pkg/util"
	"fmt"
	"io/ioutil"
	"os/exec"
)

const sshAddCommand = "ssh-add"

func AddKeys() {
	keys, err := ioutil.ReadDir(config.KeyDir())
	util.Check(err)
	
	var keyPaths []string

	for _, key := range keys {
		keyPaths = append(keyPaths, fmt.Sprintf("%v/%v", config.KeyDir(), key.Name()))
	}

	cmd := exec.Command(sshAddCommand, keyPaths...)
	bytes, err := cmd.CombinedOutput()
	util.Check(err)

	fmt.Println(string(bytes))
}
