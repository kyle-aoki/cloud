package cmd

import (
	"cloudlab/pkg/args"
	"cloudlab/pkg/resource"
	"cloudlab/pkg/util"
	"fmt"
	"os"
	"strings"
	"time"
)

func SSH() {
	names := args.CollectOrEmpty()
	allInstances := len(names) == 0

	lr := resource.NewLabResources()
	lr.Instances = resource.FindNonTerminatedInstances()

	for _, inst := range lr.Instances {
		instName := resource.FindNameTagValue(inst.Tags)
		var ip string
		if args.Flags.Private {
			ip = *inst.PrivateIpAddress
		} else {
			if inst.PublicIpAddress == nil {
				ip = *inst.PrivateIpAddress
			} else {
				ip = *inst.PublicIpAddress
			}
		}
		if !allInstances && instName != nil && util.Contains(*instName, names) {
			printSSHCommand(ip)
		}
		if allInstances {
			printSSHCommand(ip)
		}
	}
}

func printSSHCommand(ip string) {
	fmt.Printf("ssh -i %s ubuntu@%s\n", resource.KeyFilePath(), ip)
}

const SshConfigHost = `
Host $Host
    HostName $HostName
    User ubuntu
    IdentityFile $IdentityFile
    StrictHostKeyChecking no
`

func formatHost(host, hostname string) string {
	sshConfig := strings.Replace(SshConfigHost, "$Host", host, 1)
	sshConfig = strings.Replace(sshConfig, "$HostName", hostname, 1)
	sshConfig = strings.Replace(sshConfig, "$IdentityFile", resource.KeyFilePath(), 1)
	return sshConfig
}

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
