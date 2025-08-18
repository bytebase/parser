package cql

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// ParseCQL parses a CQL statement and returns the parse tree.
func ParseCQL(statement string) (antlr.Tree, error) {
	lexer := NewCqlLexer(antlr.NewInputStream(statement))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCqlParser(stream)

	// Remove default error listeners
	lexer.RemoveErrorListeners()
	p.RemoveErrorListeners()

	// Add custom error listener to capture errors
	lexerErrors := NewErrorListener()
	parserErrors := NewErrorListener()
	
	lexer.AddErrorListener(lexerErrors)
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true

	// Parse from root rule
	tree := p.Root()

	// Check for errors
	if len(lexerErrors.Errors) > 0 {
		return nil, fmt.Errorf("lexer error: %v", lexerErrors.Errors[0].Error())
	}
	if len(parserErrors.Errors) > 0 {
		return nil, fmt.Errorf("parser error: %v", parserErrors.Errors[0].Error())
	}

	return tree, nil
}