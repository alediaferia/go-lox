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
	default:
		err = s.newError(s.line, "", "Unexpected character")
	}

	return err
}

func (s *Scanner) newError(line int64, where string, message string) error {
	return fmt.Errorf("[line %d] Error: %s; %s", line, where, message)
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
