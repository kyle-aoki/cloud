# dependencies

```
go version
> 1.18
```

# install cloudlab

Build the project, then place it into a directory included in $PATH.

```
go build
mv lab ~/bin/lab

lab version
> 1.0.0
```

# syntax

```
general flags:
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
                    --type       -t   instance type (t2.micro, t2.medium...) (default t2.nano)
                    --gigabytes  -g   gigabytes of storage (integer) (default 8)

  start         <names...>            start instance(s)
  stop          <names...>            stop instance(s)

  delete        <names...>            terminate instance(s)
  open-port     <name> <ports...>     open one or more ports on an instance (all protocols)
  close-port    <name> <ports...>     close one or more ports on an instance
```