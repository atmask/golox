package lox

import "testing"

func TestScanSingleCharTokens(t *testing.T) {
	scanner := NewScanner("(){},.-+;*")
	tokens := scanner.ScanTokens()

	// expected types in order - the last one is always EOF
	expected := []TokenType{
		LEFT_PAREN, RIGHT_PAREN,
		LEFT_BRACE, RIGHT_BRACE,
		COMMA, DOT, MINUS, PLUS,
		SEMICOLON, STAR,
		EOF,
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, tok := range tokens {
		if tok.Type != expected[i] {
			t.Errorf("token[%d]: expected type %d, got %d (lexeme=%q)", i, expected[i], tok.Type, tok.Lexeme)
		}
	}
}

func TestScanEmptySource(t *testing.T) {
	scanner := NewScanner("")
	tokens := scanner.ScanTokens()

	if len(tokens) != 1 {
		t.Fatalf("expected 1 token (EOF), got %d", len(tokens))
	}
	if tokens[0].Type != EOF {
		t.Errorf("expected EOF, got %d", tokens[0].Type)
	}
}
