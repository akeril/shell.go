package builtins

import (
	"fmt"
	"os"
)

func Pwd(fdout, fderr *os.File, args []string) {
	fmt.Fprintf(fdout, "%s\n", os.Getenv("PWD"))
}
