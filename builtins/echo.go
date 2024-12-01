package builtins

import "fmt"

func Echo(args []string) error {
	n := len(args)
	for _, arg := range args[1 : n-1] {
		fmt.Printf("%v ", arg)
	}
	fmt.Println(args[n-1])
	return nil
}
