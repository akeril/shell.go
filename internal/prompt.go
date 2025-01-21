package internal

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"golang.org/x/term"
)

func ReadPrompt(engine *Trie, prompt string) (string, error) {
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
			prompt = ""
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
			doubletab = !doubletab
		}
	}

	term.Restore(int(os.Stdin.Fd()), state)
	fmt.Fprintf(os.Stdout, "\n")
	return prompt, err
}

func Autocomplete(engine *Trie, prompt string, doubletab bool) (string, string, error) {
	matches := engine.Match(prompt)
	if len(matches) == 1 {
		return prompt + matches[0] + " ", matches[0] + " ", nil
	}
	if match, ok := engine.LongestMatch(prompt); ok {
		return prompt + match, match, nil
	}
	if len(matches) == 0 || !doubletab {
		return prompt, "\a", nil
	}
	slices.Sort(matches)
	s := "\r\n"
	for _, match := range matches {
		s += prompt + match + "  "
	}
	return prompt, s, errors.New("Input terminated by autocomplete")
}
