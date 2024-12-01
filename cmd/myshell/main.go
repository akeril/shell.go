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
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic("???")
		}
		cmd = strings.TrimRight(cmd, "\n")
		fmt.Printf("%v: command not found\n", cmd)
	}
}
