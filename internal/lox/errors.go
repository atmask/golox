package lox

import "fmt"

type ScanError struct {
	Line    int
	Message string
}

func (e *ScanError) Error() string {
	return fmt.Sprintf("[line %d] Error: %s", e.Line, e.Message)
}
