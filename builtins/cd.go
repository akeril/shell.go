package builtins

import (
	"errors"
	"fmt"
	"os"
)

func Cd(args []string) error {
	path := args[1]
	if len(args) != 2 {
		return errors.New("Incorrect number of parameters")
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
