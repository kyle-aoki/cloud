package config

import (
	"cloud/pkg/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const ConfigDir = ".cloud"
const ConfigFileName = "config.json"

func FullPath() string {
	hd, err := os.UserHomeDir()
	util.Check(err)
	return fmt.Sprintf("%v/%v/%v", hd, ConfigDir, ConfigFileName)
}

var Vars ConfigFile

func init() {
	homedir, err := os.UserHomeDir()
	util.Check(err)

	configDirPath := fmt.Sprintf("%v/%v", homedir, ConfigDir)
	exists, err := fsObjectExists(configDirPath)
	util.Check(err)

	if !exists {
		err = os.Mkdir(configDirPath, os.ModePerm)
		util.Check(err)
	}

	configFilePath := fmt.Sprintf("%v/%v", configDirPath, ConfigFileName)
	exists, err = fsObjectExists(configFilePath)
	util.Check(err)

	if !exists {
		err = os.WriteFile(configFilePath, []byte(baseConfig), os.ModePerm)
		util.Check(err)
		util.Exit(configFileCreationMessage)
	}

	bytes, err := ioutil.ReadFile(configFilePath)
	util.Check(err)

	err = json.Unmarshal(bytes, &Vars)
	util.Check(err)
}

func fsObjectExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
