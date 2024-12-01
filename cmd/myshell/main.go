package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// infinite repl
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		prompt, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic("???")
		}
		args := strings.Fields(prompt)
		cmd, args := args[0], args[1:]

		if cmd == "exit" {
			os.Exit(0)
		}
		fmt.Printf("%v: command not found\n", cmd)
	}

}
