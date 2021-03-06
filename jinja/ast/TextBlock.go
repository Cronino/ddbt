package ast

import (
	"strings"
	"unicode"

	"ddbt/compilerInterface"
	"ddbt/jinja/lexer"
)

// A block which represents a simple
type TextBlock struct {
	position lexer.Position
	value    string
}

var _ AST = &TextBlock{}

func NewTextBlock(token *lexer.Token) *TextBlock {
	return &TextBlock{
		position: token.Start,
		value:    token.Value,
	}
}

func (tb *TextBlock) Position() lexer.Position {
	return tb.position
}

func (tb *TextBlock) Execute(_ compilerInterface.ExecutionContext) (*compilerInterface.Value, error) {
	return &compilerInterface.Value{StringValue: tb.value}, nil
}

func (tb *TextBlock) String() string {
	return tb.value
}

func (tb *TextBlock) TrimPrefixWhitespace() string {
	tb.value = strings.TrimLeftFunc(tb.value, unicode.IsSpace)

	return tb.value
}

func (tb *TextBlock) TrimSuffixWhitespace() string {
	tb.value = strings.TrimRightFunc(tb.value, unicode.IsSpace)

	return tb.value
}
