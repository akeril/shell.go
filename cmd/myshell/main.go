package main

import (
	"fmt"
	"os"

	"github.com/kjabin/shell.go/builtins"
	"github.com/kjabin/shell.go/internal"
)

func main() {

	cmds := append(builtins.Builtins(), internal.LocateCmdNames()...)
	trie := internal.NewTrie()
	for _, cmd := range cmds {
		trie.Insert(cmd)
	}

	// infinite repl
	prefix := ""
	for {
		fmt.Fprintf(os.Stdout, "$ %s", prefix)

		// Wait for user input
		prompt, err := internal.ReadPrompt(trie, prefix)
		if err != nil {
			prefix = prompt
			continue
		} else {
			prefix = ""
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

		if f, ok := builtins.Match(cmd); ok {
			f(fout, ferr, args)
		} else if path, ok := internal.MatchCmd(cmd); ok {
			internal.Exec(fout, ferr, cmd, path, args)
		} else {
			fmt.Fprintf(os.Stderr, "%v: command not found\n", cmd)
		}
	}
}
