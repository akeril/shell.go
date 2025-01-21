package builtins

import (
	"fmt"
	"os"

	"github.com/akeril/shell.go/internal"
)

func Type(fdout, fderr *os.File, args []string) {
	if len(args) < 2 {
		fmt.Fprintf(fderr, "Need at least one argument")
		return
	}
	for _, cmd := range args[1:] {
		if _, ok := Match(cmd); ok {
			fmt.Fprintf(fdout, "%v is a shell builtin\n", cmd)
		} else if path, ok := internal.MatchCmd(cmd); ok {
			fmt.Fprintf(fdout, "%v is %v\n", cmd, path)
		} else {
			fmt.Fprintf(fderr, "%v: not found\n", cmd)
		}
	}
}
