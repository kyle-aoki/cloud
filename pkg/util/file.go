package util

import (
	"io/ioutil"
	"os"
)

// Mode Append
func CreateEmptyFile(path string) {
	err := ioutil.WriteFile(path, []byte{}, os.ModeAppend)
	MustExec(err)
}

func CreateDir(dir string) {
	err := os.Mkdir(dir, os.ModeDir)
	MustExec(err)
}

func ObjectExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
