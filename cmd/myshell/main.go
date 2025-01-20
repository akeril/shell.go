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

	cmds := append(builtins.Builtins(), internal.LocateCmdNames()...)
	trie := internal.NewTrie()
	for _, cmd := range cmds {
		trie.Insert(cmd)
	}

	// infinite repl
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		prompt, err := ReadPrompt(trie)
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

func ReadPrompt(engine *internal.Trie) (string, error) {
	prompt := ""
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	doubletab := false
	c := make([]byte, 1)
loop:
	for {
		os.Stdin.Read(c)
		switch c[0] {
		case 3:
			fmt.Fprintf(os.Stdout, "^C")
			err = errors.New("Input terminated")
			break loop
		case 9:
			msg := ""
			prompt, msg, err = Autocomplete(engine, prompt, doubletab)
			fmt.Fprintf(os.Stdout, "%s", msg)
			if err != nil {
				break loop
			}
		case 10, 13:
			break loop
		case 127:
			if len(prompt) > 0 {
				fmt.Fprintf(os.Stdout, "\b \b")
				prompt = prompt[:len(prompt)-1]
			}
		default:
			if c[0] >= 32 {
				fmt.Fprintf(os.Stdout, "%c", c[0])
				prompt += string(c[0])
			}
		}
		if c[0] == 9 {
			doubletab = true
		}
	}

	term.Restore(int(os.Stdin.Fd()), state)
	fmt.Fprintf(os.Stdout, "\n")
	return prompt, err
}

func Autocomplete(engine *internal.Trie, prompt string, doubletab bool) (string, string, error) {
	matches := engine.Match(prompt)
	if len(matches) == 1 {
		return prompt + matches[0] + " ", matches[0] + " ", nil
	}
	if len(matches) == 0 || !doubletab {
		return prompt, "\a", nil
	}
	s := "\r\n"
	for _, match := range matches {
		s += prompt + match + "  "
	}
	return prompt, s, errors.New("Input terminated by autocomplete")
}
