package lox

import "strconv"

type Scanner struct {
	source  string
	tokens  []Token
	errors  []ScanError
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

func (s *Scanner) ScanTokens() ([]Token, []ScanError) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens, s.errors
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// Reads the current byte and then advances
// Returns the read byte
func (s *Scanner) advance() byte {
	ch := s.source[s.current]
	s.current++
	return ch

}

// match is a conditional form of advance()
// Advances iff the next byte == expected
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true

}

// peek returns the next character without advancing
// iff we are not at the end of the source
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.source[s.current]
}

// peekNest returns the next+1 character iff
// the next+1 character is not beyond the bounds of the source
func (s *Scanner) peekNext() byte {
	if (s.current + 1) >= len(s.source) {
		return '\x00'
	}
	return s.source[s.current+1]
}

// string is utility for parsing string literals when
// a leading '"' is encountered while scanning
func (s *Scanner) string() {
	for !s.isAtEnd() && s.peek() != '"' {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.errors = append(s.errors, ScanError{s.line, "Unterminated string."})
		return
	}

	// the closing '"'
	s.advance()

	literal := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, literal)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		// consume the "."
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	literal, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		s.errors = append(s.errors, ScanError{s.line, "Invalid number: " + err.Error()})
		return
	}
	s.addToken(NUMBER, literal)
}

// isDigit is a utility free function that checks if a byte
// is between the sequential literals of '0' and '9'
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// Given a TokenType and literal, append a Token to the instance slice
func (s *Scanner) addToken(t TokenType, literal any) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, Token{t, text, literal, s.line})
}

func (s *Scanner) scanToken() {
	// s.tokens = append(s.tokens, Token{})
	ch := s.advance()
	switch ch {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '*':
		s.addToken(STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case '/':
		if s.match('/') {
			for !s.isAtEnd() && s.peek() != '\n' {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
	case ' ', '\r', '\t':
		// Ignore whitespace``
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		switch {
		case isDigit(ch):
			s.number()
		default:
			s.errors = append(s.errors, ScanError{s.line, "Unexpected character: " + string(ch)})
		}
	}
}
