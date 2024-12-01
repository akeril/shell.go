package builtins

import (
	"fmt"
	"os"
)

func Pwd(args []string) error {
	fmt.Println(os.Getenv("PWD"))
	return nil
}
