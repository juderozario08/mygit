package main

import (
	"flag"
	"fmt"
)

func Commit() {
	var message string
	flag.StringVar(&message, "message", "First Commit", "Add a message for the commit latest commit")
	flag.StringVar(&message, "m", "First Commit", "Add a message for the commit latest commit")
	flag.Parse()
	fmt.Println("strFlag value is: ", message)
}
