package main

import (
	"fmt"
	"os"
)

func helpMessage() {
	data, err := os.ReadFile("./help.md")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
