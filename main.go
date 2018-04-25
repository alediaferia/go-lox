package main

import (
	"os"
)

func main() {
	p := NewPrompt(os.Stdin, os.Stdout)
	p.Run()
}
