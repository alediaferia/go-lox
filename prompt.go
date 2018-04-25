package main

import (
	"bufio"
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
	s := &Scanner{}

	for {
		_, err := w.WriteString(" > ")
		w.Flush()
		if err != nil {
			panic(err)
		}
		line, err := r.ReadString('\n')
		tokens := s.ScanTokens(line)
		println(tokens)
	}
}
