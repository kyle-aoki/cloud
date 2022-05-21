package command

import "fmt"

func SSH() {
	fmt.Println()
	fmt.Println("chmod 400 /path/to/key/file")
	fmt.Println("ssh-add /path/to/key/file")
	fmt.Println("ssh ubuntu@<public-ip-address>")
	fmt.Println()
}
