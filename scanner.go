package main

import (
	"fmt"
	"strings"
)

type Scanner struct {
	source               *strings.Reader
	start, current, line int64
	tokens               []*Token
	hasError             bool
}

func NewScanner(source string) *Scanner {
	return &Scanner{strings.NewReader(source), 0, 0, 1, make([]*Token, 0), false}
}

func (s *Scanner) ScanTokens() ([]*Token, error) {
	var err error
	for !s.atEnd() {
		s.start = s.current
		err = s.scanToken()
	}

	if err != nil {
		return s.tokens, err
	}
	s.tokens = append(s.tokens, &Token{EOF, "", nil, s.line})

	return s.tokens, nil
}

func (s *Scanner) atEnd() bool {
	return s.current >= s.source.Size()
}

func (s *Scanner) scanToken() error {
	ch, err := s.advance()
	if err != nil {
		return err
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
	case '!':
		if s.match('=') {
			s.addToken(BangEqual)
		} else {
			s.addToken(Bang)
		}
	case '=':
		if s.match('=') {
			s.addToken(EqualEqual)
		} else {
			s.addToken(Less)
		}
	case '<':
		if s.match('=') {
			s.addToken(LessEqual)
		} else {
			s.addToken(Less)
		}
	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual)
		} else {
			s.addToken(Greater)
		}
	default:
		err = s.newError(s.line, "", "Unexpected character")
	}

	return err
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
	text := make([]byte, s.current-s.start)
	_, err := s.source.ReadAt(text, s.start)
	if err != nil {
		return err
	}

	s.tokens = append(s.tokens, &Token{t, string(text), literal, s.line})

	return nil
}

func (s *Scanner) match(ch byte) bool {
	if s.atEnd() {
		return false
	}

	b := make([]byte, 1)
	_, err := s.source.ReadAt(b, s.current)
	if err != nil {
		panic(err)
	}

	if b[0] != ch {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) newError(line int64, where string, message string) error {
	return fmt.Errorf("[line %d] Error: %s; %s", line, where, message)
}
