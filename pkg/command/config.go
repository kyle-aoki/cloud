package command

import (
	"cloud/pkg/config"
	"cloud/pkg/util"
	"encoding/json"
	"fmt"
)

func ShowConfig() {
	mijson, err := json.MarshalIndent(config.Vars, "", "  ")
	util.MustExec(err)
	fmt.Print(string(mijson))
}
