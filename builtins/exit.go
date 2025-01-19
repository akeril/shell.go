package builtins

import "os"

func Exit(fdout, fderr *os.File, args []string) {
	os.Exit(0)
}
