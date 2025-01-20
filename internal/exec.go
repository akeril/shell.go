package internal

import (
	"os"
	"syscall"
)

func Exec(fdout, fderr *os.File, cmd string, path string, args []string) error {
	pid, err := syscall.ForkExec(path, args, &syscall.ProcAttr{
		Env: os.Environ(),
		Files: []uintptr{
			uintptr(syscall.Stdin),
			fdout.Fd(),
			fderr.Fd(),
		},
	})
	if err != nil {
		return err
	}
	_, err = syscall.Wait4(pid, nil, 0, nil)
	return err
}
