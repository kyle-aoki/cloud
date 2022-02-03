package config

import (
	"cloud/pkg/util"
	"fmt"
	"os"
)

const keys = "keys"

func KeyDir() string {
	hd, err := os.UserHomeDir()
	util.Check(err)
	keyDir := fmt.Sprintf("%v/%v/%v", hd, ConfigDirName, keys)
	if _, err := os.Stat(keyDir); os.IsNotExist(err) {
		err := os.MkdirAll(keyDir, os.ModePerm)
		util.Check(err)
	}
	return keyDir
}
