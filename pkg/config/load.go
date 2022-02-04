package config

import (
	"cloud/pkg/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFileName = ".cloud"
var Vars ConfigFile

func Load() {
	hd, err := os.UserHomeDir()
	util.MustExec(err)

	ConfigFilePath := fmt.Sprintf("%v/%v", hd, configFileName)

	exists, err := FsObjectExists(ConfigFilePath)
	util.MustExec(err)

	if !exists {
		err := os.WriteFile(ConfigFilePath, []byte(baseConfig), os.ModePerm)
		util.MustExec(err)
		fmt.Print(configFileCreationMessage)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(ConfigFilePath)
	util.MustExec(err)

	err = json.Unmarshal(bytes, &Vars)
	util.MustExec(err)
}

func FsObjectExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
