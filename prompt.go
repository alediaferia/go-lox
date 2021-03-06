package main

import (
	"bufio"
	"fmt"
	"io"
)

// Prompt  is a simple class that can be used to implement
// a simple REPL console application
type Prompt struct {
	reader io.Reader
	writer io.Writer
}

// NewPrompt returns an instance of Prompt
func NewPrompt(reader io.Reader, writer io.Writer) *Prompt {
	return &Prompt{reader, writer}
}

func (p *Prompt) Run() {
	r := bufio.NewReader(p.reader)
	w := bufio.NewWriter(p.writer)

	for {
		_, err := w.WriteString(" > ")
		w.Flush()
		if err != nil {
			panic(err)
		}
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		s := NewScanner(string(line[:len(line)-1]))
		tokens, err := s.ScanTokens()
		if err != nil {
			fmt.Println(err)
		} else {
			for _, t := range tokens {
				fmt.Println(t)
			}
		}
	}
}
