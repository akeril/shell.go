package builtins

import (
	"errors"
	"fmt"

	"github.com/kjabin/shell.go/internal"
)

func Type(args []string) error {
	if len(args) < 2 {
		return errors.New("Need at least one argument")
	}
	for _, cmd := range args[1:] {
		if _, ok := Match(cmd); ok {
			fmt.Printf("%v is a shell builtin\n", cmd)
		} else if path, ok := internal.MatchCmd(cmd); ok {
			fmt.Printf("%v is %v\n", cmd, path)
		} else {
			fmt.Printf("%v: not found\n", cmd)
		}
	}
	return nil
}
