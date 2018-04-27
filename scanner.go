package main

import (
	"strings"
)

type Scanner struct {
	source               *strings.Reader
	start, current, line int64
	tokens               []*Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{strings.NewReader(source), 0, 0, 1, make([]*Token, 0)}
}

func (s *Scanner) ScanTokens() []*Token {
	for !s.atEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, &Token{EOF, "", nil, s.line})

	return s.tokens
}

func (s *Scanner) atEnd() bool {
	return s.current >= s.source.Size()
}

func (s *Scanner) scanToken() {
	ch, err := s.advance()
	if err != nil {
		panic(err)
	}

	switch ch {
	case '(':
		s.addToken(LeftParen)
	case ')':
		s.addToken(RightParen)
	case '{':
		s.addToken(LeftBrace)
	case '}':
		s.addToken(RightBrace)
	case ',':
		s.addToken(Comma)
	case '.':
		s.addToken(Dot)
	case '-':
		s.addToken(Minus)
	case '+':
		s.addToken(Plus)
	case ';':
		s.addToken(Semicolon)
	case '*':
		s.addToken(Star)
	}
}

func (s *Scanner) advance() (rune, error) {
	ch, size, err := s.source.ReadRune()
	s.current += int64(size)
	return ch, err
}

func (s *Scanner) addToken(t TokenType) {
	s.addTokenAndLiteral(t, nil)
}

func (s *Scanner) addTokenAndLiteral(t TokenType, literal interface{}) error {
	s.source.Seek(s.start, 0)
	var text string
	for i := int64(0); i < s.current-s.start; {
		ch, size, err := s.source.ReadRune()
		if err != nil {
			return err
		}
		text = text + string(ch)
		i += int64(size)
	}

	s.tokens = append(s.tokens, &Token{t, text, literal, s.line})

	return nil
}
