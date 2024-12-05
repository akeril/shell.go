package internal

import (
	"strings"
)

func SplitArgs(s string) []string {
	var args []string
	s = strings.Trim(s, "\r\n")
	buf := ""
	isSingleQuoted := false
	isDoubleQuoted := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '\\' && !isSingleQuoted && !isDoubleQuoted {
			if i+1 < len(s) {
				buf += string(s[i+1])
				i++
			}
		} else if c == '\\' && isDoubleQuoted {
			if i+1 < len(s) && (s[i+1] == '$' || s[i+1] == '\\' || s[i+1] == '"') {
				buf += string(s[i+1])
				i++
			} else {
				buf += "\\"
			}
		} else if c == '\'' && !isDoubleQuoted {
			isSingleQuoted = !isSingleQuoted
		} else if c == '"' && !isSingleQuoted {
			isDoubleQuoted = !isDoubleQuoted
		} else if c == ' ' && !isSingleQuoted && !isDoubleQuoted {
			appendArg(&args, &buf)
		} else {
			buf += string(c)
		}
	}
	appendArg(&args, &buf)
	return args
}

func appendArg(args *[]string, buf *string) {
	if *buf != "" {
		*args = append(*args, *buf)
	}
	*buf = ""
}
