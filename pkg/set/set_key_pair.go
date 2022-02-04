package set

import (
	"cloud/pkg/args"
	"cloud/pkg/config"
	"cloud/pkg/get"
	"cloud/pkg/util"
	"encoding/json"
	"io/ioutil"
	"os"
)

// cloud set keypair <keypair-name> <node-config-1> <node-config-2> ...
func SetKeyPair() {
	keypairName := args.Poll()
	nodeConfigNames := args.Collect()

	confirmKeyPairExists(keypairName)

	for i := range config.Vars.NodeConfigs {
		if contains(nodeConfigNames, config.Vars.NodeConfigs[i].Name) {
			config.Vars.NodeConfigs[i].KeyPair = keypairName
		}
	}
	configJson, err := json.MarshalIndent(&config.Vars, "", "  ")
	util.Check(err)
	ioutil.WriteFile(config.FullPath(), configJson, os.ModePerm)
}

func contains(nodeConfigNames []string, nodeConfig string) bool {
	for i := range nodeConfigNames {
		if nodeConfigNames[i] == nodeConfig {
			return true
		}
	}
	return false
}

func confirmKeyPairExists(keyPairName string) bool {
	dkp := get.GetKeyPairs()
	for _, kp := range dkp.KeyPairs {
		if *kp.KeyName == keyPairName {
			return true
		}
	}
	panic("key pair does not exist")
}
