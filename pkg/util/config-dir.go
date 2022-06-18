package util

import "fmt"

const ConfigName = ".cloudlab"

func ConfigDir() string {
	return fmt.Sprintf("%s/%s", HomeDir(), ConfigName)
}
