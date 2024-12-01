package builtins

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/internal"
)

func Type(args []string) error {
	cmd := args[0]
	if _, ok := Match(cmd); ok {
		fmt.Printf("%v is a shell builtin\n", cmd)
	} else if path, ok := internal.MatchExecutable(cmd); ok {
		fmt.Printf("%v is %v\n", cmd, path)
	} else {
		fmt.Printf("%v: not found\n", cmd)
	}
	return nil
}
