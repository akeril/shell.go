package internal

import (
	"strings"
)

/*
what the funck
*/
func SplitArgs(s string) []string {
	var args []string
	s = strings.Trim(s, "\r\n")
	buf := ""
	isQuoted := false
	for _, c := range s {
		if c == ' ' {
			if isQuoted {
				buf += string(c)
			} else if buf != "" {
				args = append(args, buf)
				buf = ""
			}
		} else if c == '\'' {
			isQuoted = !isQuoted
			if buf != "" {
				args = append(args, buf)
				buf = ""
			}
		} else {
			buf += string(c)
		}
	}
	if buf != "" {
		args = append(args, buf)
	}
	return args
}
