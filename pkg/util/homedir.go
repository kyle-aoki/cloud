package util

import "os"

func HomeDir() string {
	hd, err := os.UserHomeDir()
	MustExec(err)
	return hd
}
