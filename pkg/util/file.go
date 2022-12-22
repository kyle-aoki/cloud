package util

import (
	"log"
	"os"
)

func CreateDir(dir string) {
	log.Println("creating dir", dir)
	err := os.Mkdir(dir, 0777)
	Check(err)
}

func ObjectExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println(path, "does not exist")
		return false
	}
	log.Println(path, "does exist")
	return true
}
