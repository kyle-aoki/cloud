package command

import (
	"cloud/pkg/args"
	"cloud/pkg/defaults"
	"strconv"
)

func OpenPort() {
	port := args.Poll()
	instnaceName := args.Poll()

	cldo := defaults.Start()

	sg := cldo.GetSecurityGroupIdByNameOrNil(port)

	if sg == nil {
		portInt64, err := strconv.ParseInt(port, 10, 32)
		if err != nil {
			panic("invalid port")
		}
		portInt := int(portInt64)

		if portInt > 65535 || portInt < 1 {
			panic("invalid port")
		}
		cldo.CreateSecurityGroup(port, portInt)
	}

	instance := cldo.GetInstanceByName(instnaceName)
	securityGroup := cldo.GetSecurityGroupIdByNameOrPanic(port)

	cldo.AssignSecurityGroup(instance, securityGroup)
}
