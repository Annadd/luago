package main

import (
	. "luago/api"
	"luago/state"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		ls := state.New()
		ls.OpenLibs()
		if err := ls.LoadFile(os.Args[1]); err == LUA_OK {
			ls.Call(0, -1)
		}
	}
}
