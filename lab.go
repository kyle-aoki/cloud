package main

import (
  "cloudlab/pkg/amazon"
  "cloudlab/pkg/args"
  "cloudlab/pkg/cmd"
  "cloudlab/pkg/util"
  "fmt"
  "io"
  "log"
  "os"
)

const cloudlabVersion = "1.1.0"

func main() {
  defer util.MainRecover()
  args.ParseProgramInput()

  if args.Flags.PrintVersion {
    fmt.Println(cloudlabVersion)
    os.Exit(0)
  }
  if args.Flags.Verbose {
    log.SetOutput(os.Stdout)
  } else {
    log.SetOutput(io.Discard)
  }
  if args.Flags.ShowHelp {
    fmt.Print(helpText)
    os.Exit(0)
  }

  if command, found := syntax[args.PollOrEmpty()]; found {
    amazon.InitEC2Client()
    command()
  } else {
    fmt.Print(helpText)
  }
}

var syntax = map[string]func(){
  "info":       cmd.Info,
  "init":       cmd.InitializeCloudLabResources,
  "destroy":    cmd.DestroyCloudLabResources,
  "list":       cmd.ListInstances,
  "watch":      cmd.Watch,
  "run":        cmd.Run,
  "start":      cmd.StartInstance,
  "stop":       cmd.StopInstance,
  "ssh":        cmd.SSH,
  "delete":     cmd.DeleteInstances,
  "open-port":  cmd.OpenPorts,
  "close-port": cmd.ClosePorts,
}

const helpText = `general flags:
  -v, --version    print version
  -h, --help       print this help text
  --verbose        verbose logging
commands:
  version       print version
  info          print cloudlab resource info

  init          create base cloudlab resources (vpc, subnets, etc.)
                (base resources cost nothing)
  destroy       destroy base cloudlab resrouces
                (must terminate all instances first)

  list          list active instances
                    --all    -a       show terminated instances
                    --quiet  -q       print names only

  watch         run 'lab list' continuously
                    --all    -a       show terminated instances

  run           run a new instance
                    --name       -n   instance name
                    --private    -p   create instance in private subnet
                    --type       -t   instance type (t2.micro, t2.medium, ...) (default t2.nano)
                    --gigabytes  -g   gigabytes of storage (integer) (default 8)

  start         <names...>            start instance(s)
  stop          <names...>            stop instance(s)
  ssh           <names...>            print out SSH command(s)
  delete        <names...>            terminate instance(s)
  open-port     <name> <ports...>     open one or more ports on an instance (all protocols)
  close-port    <name> <ports...>     close one or more ports on an instance
`
