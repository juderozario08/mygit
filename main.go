package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Use one of the following methods")
		helpMessage()
		os.Exit(ERROR)
	}
	args := os.Args[1:]
	switch a := args[0]; a {
	case "add":
		Add()
	case "cat-file":
		CatFile()
	case "check-ignore":
		CheckIgnore()
	case "checkout":
		Checkout()
	case "commit":
		Commit()
	case "hash-object":
		HashObject()
	case "init":
		Init()
	case "log":
		Log()
	case "ls-files":
		LsFiles()
	case "ls-tree":
		LsTree()
	case "rev-parse":
		RevParse()
	case "rm":
		Rm()
	case "show-ref":
		ShowRef()
	case "status":
		Status()
	case "tag":
		Tag()
	default:
		helpMessage()
	}
	os.Exit(SUCCESS)
}
