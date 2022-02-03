package show

import (
	"cloud/pkg/config"
	"cloud/pkg/util"
	"fmt"
	"io/ioutil"
)

func ShowKeyPairs() {
	keys, err := ioutil.ReadDir(config.KeyDir())
	util.Check(err)
	for _, key := range keys {
		kpath := fmt.Sprintf("%v/%v", config.KeyDir(), key.Name())
		fmt.Println(kpath)
	}
}
