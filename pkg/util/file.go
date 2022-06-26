package util

import (
	"io/ioutil"
	"log"
	"os"
)

// Mode Append
func CreateEmptyFile(path string) {
	log.Println("creating empty file", path)
	err := ioutil.WriteFile(path, []byte{}, 0777)
	MustExec(err)
}

func CreateDir(dir string) {
	log.Println("creating dir", dir)
	err := os.Mkdir(dir, 0777)
	MustExec(err)
}

func ObjectExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println(path, "does not exist")
		return false
	}
	log.Println(path, "does exist")
	return true
}

func DeleteFile(path string) {
	err := os.Remove(path)
	MustExec(err)
}
