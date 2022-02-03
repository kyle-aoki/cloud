package show

import (
	"cloud/pkg/config"
	"cloud/pkg/util"
	"fmt"
	"io/ioutil"
)

func ShowConfig() {
	contents, err := ioutil.ReadFile(config.FullPath())
	util.Check(err)
	fmt.Println(string(contents))
}
