package internal

import (
	"errors"
	"os"
	"slices"
)

func RedirectStream(args []string, ops []string) ([]string, *os.File, error) {
	newArgs := make([]string, 0, len(args))
	streams := make([]Redirect, 0)
	i := 0
	for i+1 < len(args) {
		if !slices.Contains(ops, args[i]) {
			newArgs = append(newArgs, args[i])
			i += 1
		} else {
			streams = append(streams, Redirect{args[i], args[i+1]})
			i += 2
		}
	}
	if i < len(args) {
		newArgs = append(newArgs, args[i])
	}
	if len(streams) == 0 {
		return newArgs, os.Stderr, nil
	}
	f, err := streams[len(streams)-1].Open()
	return newArgs, f, err
}

type Redirect struct {
	kind string
	file string
}

func (r Redirect) Open() (*os.File, error) {
	switch r.kind {
	case "1>", ">", "2>":
		return os.Create(r.file)
	case "1>>", ">>", "2>>":
		return os.OpenFile(r.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	return nil, errors.New("Invalid kind")
}
