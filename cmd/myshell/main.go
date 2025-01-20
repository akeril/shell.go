package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kjabin/shell.go/builtins"
	"github.com/kjabin/shell.go/internal"
	"golang.org/x/term"
)

func main() {

	// infinite repl
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		prompt, err := ReadPrompt()
		if err != nil {
			continue
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

func ReadPrompt() (string, error) {
	prompt := ""
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	c := make([]byte, 1)
loop:
	for {
		os.Stdin.Read(c)

		switch c[0] {
		case 3:
			fmt.Fprintf(os.Stdout, "^C")
			err = errors.New("Input terminated")
			break loop
		case 10, 13:
			break loop
		case 127:
			if len(prompt) > 0 {
				fmt.Fprintf(os.Stdout, "\b \b")
				prompt = prompt[:len(prompt)-1]
			}
		default:
			fmt.Fprintf(os.Stdout, "%c", c[0])
			prompt += string(c[0])
		}
	}

	term.Restore(int(os.Stdin.Fd()), state)
	fmt.Fprintf(os.Stdout, "\n")
	return prompt, err
}
