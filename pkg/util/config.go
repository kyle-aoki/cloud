package util

import (
	"fmt"
	"os"
)

const ConfigDirName = ".cloudlab"

func HomeDir() string {
	hd, err := os.UserHomeDir()
	Check(err)
	return hd
}

func ConfigDir() string {
	return fmt.Sprintf("%s/%s", HomeDir(), ConfigDirName)
}
