package builtins

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Cd(args []string) error {
	path := args[1]
	if len(args) != 2 {
		return errors.New("Incorrect number of parameters")
	}
	if path[0] == '~' {
		path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	}
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("cd: %v: No such file or directory\n", path)
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("cd: %v: No such file or directory\n", path)
		return err
	}
	err = os.Setenv("PWD", dir)
	return err
}
