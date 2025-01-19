package builtins

import "os"

type Command func(*os.File, *os.File, []string)

func Match(cmd string) (Command, bool) {
	mp := map[string]Command{
		"echo": Echo,
		"exit": Exit,
		"type": Type,
		"pwd":  Pwd,
		"cd":   Cd,
	}
	f, ok := mp[cmd]
	return f, ok
}
