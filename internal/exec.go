package internal

import (
	"os"
	"syscall"
)

func Exec(cmd string, path string, args []string) error {
	args[0] = path
	pid, err := syscall.ForkExec(path, args, &syscall.ProcAttr{
		Env: os.Environ(),
		Files: []uintptr{
			uintptr(syscall.Stdin),
			uintptr(syscall.Stdout),
			uintptr(syscall.Stderr),
		},
	})
	if err != nil {
		return err
	}
	_, err = syscall.Wait4(pid, nil, 0, nil)
	return err
}
