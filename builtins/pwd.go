package builtins

import (
	"fmt"
	"os"
)

func Pwd(fdout, fderr *os.File, args []string) {
	fmt.Fprintf(fdout, "%s", os.Getenv("PWD"))
}
