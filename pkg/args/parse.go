package args

import (
	"cloudlab/pkg/util"
	"strings"
)

func parseCloudlabFlags(args []string) (*CloudlabFlags, []string) {
	nonFlagArgs := []string{}
	flags := cloudlabFlagDefaults()
	for i := 0; i < len(args); i++ {
		// -------------------------------------------------------------------------
		// ---- valueless flags ----------------------------------------------------
		// -------------------------------------------------------------------------
		if args[i] == "-v" || args[i] == "--version" {
			flags.PrintVersion = true
			continue
		}
		if args[i] == "--verbose" {
			flags.Verbose = true
			continue
		}
		if args[i] == "-q" || args[i] == "--quiet" {
			flags.Quiet = true
			continue
		}
		if args[i] == "-p" || args[i] == "--private" {
			flags.Private = true
			continue
		}
		if args[i] == "-a" || args[i] == "--all" {
			flags.ShowTerminated = true
			continue
		}
		if args[i] == "-h" || args[i] == "--help" {
			flags.ShowHelp = true
			continue
		}
		// -------------------------------------------------------------------------
		// ---- instance type ------------------------------------------------------
		// -------------------------------------------------------------------------
		if args[i] == "-t" {
			if isLastArg(args, i) {
				panic("no argument supplied for instance type. try -t <type> or --type=<type>")
			}
			flags.InstanceType = args[i+1]
			i += 1
			continue
		}
		if strings.Contains(args[i], "--type=") {
			typeParts := strings.SplitN(args[i], "=", 2)
			if len(typeParts) < 2 {
				panic("invalid instance type. try -t <type> or --type=<type>")
			}
			flags.InstanceType = typeParts[1]
			continue
		}
		// -------------------------------------------------------------------------
		// ---- gigabytes ----------------------------------------------------------
		// -------------------------------------------------------------------------
		if args[i] == "-g" {
			if isLastArg(args, i) {
				panic("no argument supplied for instance gigabytes. try -g <gigs> or --gigabytes=<gigs>")
			}
			flags.Gigabytes = args[i+1]
			i += 1
			continue
		}
		if strings.Contains(args[i], "--gigabytes=") {
			gigParts := strings.SplitN(args[i], "=", 2)
			if len(gigParts) < 2 {
				panic("invalid gigabyte value. try -g <gigs> or --gigabytes=<gigs>")
			}
			flags.Gigabytes = gigParts[1]
			continue
		}
		// -------------------------------------------------------------------------
		// ---- instance name ------------------------------------------------------
		// -------------------------------------------------------------------------
		if args[i] == "-n" {
			if isLastArg(args, i) {
				panic("no argument supplied for instance name. try -n <name> or --name=<name>")
			}
			flags.InstanceName = util.StrPtr(args[i+1])
			i += 1
			continue
		}
		if strings.Contains(args[i], "--name=") {
			nameParts := strings.SplitN(args[i], "=", 2)
			if len(nameParts) < 2 {
				panic("invalid instance name. try -n <name> or --name=<name>")
			}
			flags.InstanceName = util.StrPtr(nameParts[1])
			continue
		}
		// -------------------------------------------------------------------------
		// ---- no args found ------------------------------------------------------
		// -------------------------------------------------------------------------
		nonFlagArgs = append(nonFlagArgs, args[i])
	}
	return flags, nonFlagArgs
}
