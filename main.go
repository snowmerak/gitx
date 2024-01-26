package main

import (
	"fmt"
)

func main() {
	g, err := NewGit(".", nil)
	if err != nil {
		panic(err)
	}

	st, err := g.Status()
	if err != nil {
		panic(err)
	}

	for k, v := range st {
		fmt.Println(k, byte(v.Staging), byte(v.Worktree))
	}
}
