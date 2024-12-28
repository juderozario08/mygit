package main

import (
	"os"
)

func Init(path string, force bool) {
	gitdir := path + "/.git"
	info, err := os.Stat(gitdir)
	if err != nil {
		panic(err)
	}
	if !(force || info.IsDir()) {
		panic("Not a git repository " + gitdir)
	}
}
