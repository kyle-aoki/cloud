package config

import (
	"cloud/pkg/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFileName = ".cloud"

var Vars ConfigVars

func ConfigFilePath() string {
	homedir, err := os.UserHomeDir()
	util.MustExec(err)

	return fmt.Sprintf("%v/%v", homedir, configFileName)
}

func Load() {
	if !util.FileExists(ConfigFilePath()) {
		err := ioutil.WriteFile(ConfigFilePath(), []byte(DEFAULT_CONFIG), os.ModePerm)
		util.MustExec(err)
		fmt.Println(CONFIG_FILE_CREATED)
	}
	configFile, err := ioutil.ReadFile(ConfigFilePath())
	util.MustExec(err)

	err = json.Unmarshal(configFile, &Vars)
	util.MustExec(err)
}
