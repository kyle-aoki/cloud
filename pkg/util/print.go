package util

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var W *tabwriter.Writer

func init() { W = tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0) }

func Tab(a ...any) { fmt.Fprintln(W, a...) }

func ExecPrint() { W.Flush() }
