package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/tabwriter"
)

var tWriter *tabwriter.Writer

func init() { tWriter = tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0) }

func Tab(a ...any) { fmt.Fprintln(tWriter, a...) }

func TabPrint(strs ...string) {
	var tabstr string
	for i := 0; i < len(strs); i++ {
		tabstr += strs[i]
		if i != len(strs)-1 {
			tabstr += "\t"
		}
	}
	Tab(tabstr)
}

func ExecPrint() { tWriter.Flush() }

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}
