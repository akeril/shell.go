package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kjabin/shell.go/builtins"
	"github.com/kjabin/shell.go/internal"
)

func main() {

	// infinite repl
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		prompt, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}
		args := internal.SplitArgs(prompt)
		cmd := args[0]

		args, fout, err := internal.RedirectStream(args, []string{">", "1>", "1>>", ">>"})
		if err != nil {
			fmt.Printf("%s: %s\n", cmd, err)
			continue
		}

		args, ferr, err := internal.RedirectStream(args, []string{"2>", "2>>"})
		if err != nil {
			fmt.Printf("%s: %s\n", cmd, err)
			continue
		}

		fmt.Println(args, fout, ferr)

		if f, ok := builtins.Match(cmd); ok {
			f(fout, ferr, args)
		} else if path, ok := internal.MatchCmd(cmd); ok {
			internal.Exec(fout, ferr, cmd, path, args)
		} else {
			fmt.Fprintf(os.Stderr, "%v: command not found\n", cmd)
		}
	}

}
