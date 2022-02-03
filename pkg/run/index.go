package run

import (
	"cloud/pkg/args"
	"cloud/pkg/create"
	"cloud/pkg/delete"
	"cloud/pkg/get"
	"cloud/pkg/help"
	"cloud/pkg/show"
)

func Program() {
	defer handlePanic()

	switch args.Poll() {
	case "create":
		switch args.Poll() {
		case "node":
			create.CreateNode()
		case "keypair":
			create.CreateKeyPair()
		}
	case "show":
		switch args.Poll() {
		case "config":
			show.ShowConfig()
		}
	case "get":
		switch args.Poll() {
		case "nodes":
			get.PrintNodes()
		}
	case "delete":
		switch args.Poll() {
		case "node":
			delete.DeleteNodes()
		case "nodes":
			delete.DeleteNodes()
		}
	default:
		help.Print()
	}
}
