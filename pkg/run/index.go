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

	f := args.Poll

	switch f() {
	case "create": switch f() {
		case "node": 				create.CreateNode()
		case "keypair": 			create.CreateKeyPair()
		}
	case "show": switch f() {
		case "config":				show.ShowConfig()
		}
	case "get": switch f() {
		case "nodes":				get.PrintNodes()
		case "keypairs":			get.PrintKeyPairs()
		}
	case "delete":switch f() {
		case "node":				delete.DeleteNodes()
		case "nodes":				delete.DeleteNodes()
		case "keypair":				delete.DeleteKeyPairs()
		case "keypairs":			delete.DeleteKeyPairs()
		}
	default:
		help.Print()
	}
}
