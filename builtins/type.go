package builtins

import "fmt"

func Type(args []string) error {
	cmd := args[0]
	if _, ok := Match(cmd); ok {
		fmt.Printf("%v is a shell builtin\n", cmd)
	} else {
		fmt.Printf("%v: not found\n", cmd)
	}
	return nil
}
