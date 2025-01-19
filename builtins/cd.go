package builtins

import (
	"fmt"
	"os"
	"strings"
)

func Cd(fdout, fderr *os.File, args []string) {
	path := args[1]
	if len(args) != 2 {
		fmt.Fprintf(fderr, "Incorrect number of parameters")
		return
	}
	if path[0] == '~' {
		path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	}
	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintf(fderr, "cd: %v: No such file or directory\n", path)
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(fderr, "cd: %v: No such file or directory\n", path)
		return
	}
	err = os.Setenv("PWD", dir)
	if err != nil {
		fmt.Fprintf(fderr, "cd: %v", err)
	}
}
