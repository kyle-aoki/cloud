package cmd

import (
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"os"
	"strings"
	"time"
)

const SshConfigHost = `
Host $Host
    HostName $HostName
    User ubuntu
    IdentityFile $IdentityFile
    StrictHostKeyChecking no
`

// Host i1
//     HostName 3.16.38.82
//     User ubuntu
//     IdentityFile /Users/kyle/.cloudlab/key.pem
//     StrictHostKeyChecking no
func formatHost(host, hostname string) string {
	sshConfig := strings.Replace(SshConfigHost, "$Host", host, 1)
	sshConfig = strings.Replace(sshConfig, "$HostName", hostname, 1)
	sshConfig = strings.Replace(sshConfig, "$IdentityFile", resource.KeyFilePath(), 1)
	return sshConfig
}

// /Users/kyle/.cloudlab/key.pem
func sshConfigFile() string {
	home := util.Must(os.UserHomeDir())
	return home + "/.ssh/config"
}

func writeInstanceToSshConfig(instName, ipAddress string) string {
	formattedHost := formatHost(instName, ipAddress)
	f := util.Must(os.OpenFile(sshConfigFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644))
	defer f.Close()
	util.Must(f.Write([]byte(formattedHost)))
	return formattedHost
}

func addInstanceToSshConfig(name string, lr *resource.LabResources) string {
	time.Sleep(5 * time.Second)
	const maxCount = 120
	var count int
	var ip *string
	for ip == nil && count <= maxCount {
		time.Sleep(2 * time.Second)
		lr.Instances = resource.FindNonTerminatedInstances()
		inst := resource.FindInstanceByName(name)
		if resource.InstanceInPrivateSubnet(inst, lr) {
			ip = resource.FindInstanceByName(name).PrivateIpAddress
		} else {
			ip = resource.FindInstanceByName(name).PublicIpAddress
		}
		count++
	}
	if count > maxCount && ip == nil {
		fmt.Printf("could not add instance '%s' to %s\n", name, sshConfigFile())
		os.Exit(0)
	}
	hostConfig := writeInstanceToSshConfig(name, *ip)
	return hostConfig
}

func RemoveInstanceFromSshConfig(instName, ipAddress *string) {
	if instName == nil || ipAddress == nil {
		return
	}
	sshConfig := formatHost(*instName, *ipAddress)
	sshConfig = strings.Trim(sshConfig, "\n")
	sshConfigFileContent := string(util.Must(os.ReadFile(sshConfigFile())))
	sshConfigFileContent = strings.Replace(string(sshConfigFileContent), sshConfig, "", -1)
	sshConfigFileContent = strings.Trim(sshConfigFileContent, "\n")
	util.MustExec(os.WriteFile(sshConfigFile(), []byte(sshConfigFileContent), os.ModePerm))
}
