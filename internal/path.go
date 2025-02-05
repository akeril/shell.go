package internal

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func MatchCmd(cmd string) (string, bool) {
	executables := LocateCmd()
	for _, executable := range executables {
		if path.Base(executable) == cmd {
			return executable, true
		}
	}
	return "", false
}

func LocateCmd() []string {
	paths := strings.Split(os.Getenv("PATH"), ":")
	executables := make([]string, 0)

	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, file := range files {
			executables = append(executables, filepath.Join(path, file.Name()))
		}
	}

	return executables
}

func LocateCmdNames() []string {
	paths := strings.Split(os.Getenv("PATH"), ":")
	executables := make([]string, 0)

	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, file := range files {
			executables = append(executables, file.Name())
		}
	}

	return executables
}
