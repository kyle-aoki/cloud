package util

import (
	"fmt"
	"os"
)

const ConfigDirName = ".cloudlab"

func HomeDir() string {
	hd, err := os.UserHomeDir()
	MustExec(err)
	return hd
}

func ConfigDir() string {
	return fmt.Sprintf("%s/%s", HomeDir(), ConfigDirName)
}
