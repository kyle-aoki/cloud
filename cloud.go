package main

import (
	"cloud/pkg/args"
	"cloud/pkg/create"
	"cloud/pkg/delete"
	"cloud/pkg/get"
	"cloud/pkg/help"
	"cloud/pkg/run"
	"cloud/pkg/set"
	"cloud/pkg/show"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	f := args.Poll

	switch f() {
	case "list":
		get.PrintNodes()

	case "create":
		switch f() {
		case "node":
			create.CreateNode()
		case "nodes":
			create.CreateNode()
		case "keypair":
			create.CreateKeyPair()
		case "keypairs":
			create.CreateKeyPair()
		default:
			help.Print()
		}

	case "show":
		switch f() {
		case "config":
			show.ShowConfig()
		case "keypairs":
			show.ShowKeyPairs()
		default:
			help.Print()
		}

	case "get":
		switch f() {
		case "nodes":
			get.PrintNodes()
		case "keypairs":
			get.PrintKeyPairs()
		default:
			help.Print()
		}

	case "delete":
		switch f() {
		case "node":
			delete.DeleteNodes()
		case "nodes":
			delete.DeleteNodes()
		case "keypair":
			delete.DeleteKeyPairs()
		case "keypairs":
			delete.DeleteKeyPairs()
		default:
			help.Print()
		}

	case "add":
		switch f() {
		case "keys":
			run.AddKeys()
		}
	case "set":
		switch f() {
		case "keypair":
			set.SetKeyPair()
		case "keypairs":
			set.SetKeyPair()
		}

	default:
		help.Print()
	}
}
