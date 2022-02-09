package command

import (
	"cloud/pkg/config"
	"cloud/pkg/util"
	"fmt"
	"io/ioutil"
)

func ShowConfig() {
	content, err := ioutil.ReadFile(config.ConfigFilePath())
	util.MustExec(err)
	fmt.Print(string(content))
}
