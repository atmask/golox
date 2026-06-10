package lox

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
// iff we are not at the end of the line
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.source[s.current]
}

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
		s.errors = append(s.errors, ScanError{s.line, "Unexpected character: " + string(ch)})
	}
}
