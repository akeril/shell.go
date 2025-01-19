package builtins

import (
	"fmt"
	"os"
)

func Echo(fdout, fderr *os.File, args []string) {
	n := len(args)
	for _, arg := range args[1 : n-1] {
		fmt.Fprintf(fdout, "%v ", arg)
	}
	fmt.Fprintf(fdout, "%s\n", args[n-1])
}
