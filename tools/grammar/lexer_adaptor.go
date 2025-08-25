/*
 * This file is translated from https://github.com/antlr/grammars-v4/blob/master/antlr/antlr4/TypeScript/LexerAdaptor.ts.
 */
package grammar

import (
	"github.com/antlr4-go/antlr/v4"
)

// LexerAdaptor extends the ANTLR lexer to handle context-sensitive parsing
// of ANTLR v4 grammar files. This is needed because ANTLR grammar syntax
// has ambiguous constructs like [abc] (charset) vs [action] that require
// context to distinguish.
type LexerAdaptor struct {
	*antlr.BaseLexer

	// Constants for special construct types
	PREQUEL_CONSTRUCT int
	OPTIONS_CONSTRUCT int

	/**
	* Track whether we are inside of a rule and whether it is lexical parser. _currentRuleType==Token.INVALID_TYPE
	* means that we are outside of a rule. At the first sign of a rule name reference and _currentRuleType
	* ==invalid, we can assume that we are starting a parser rule. Similarly, seeing a token reference when not
	* already in rule means starting a token rule. The terminating ';' of a rule, flips this back to invalid type.
	*
	* This is not perfect logic but works. For example, "grammar T;" means that we start and stop a lexical rule
	* for the "T;". Dangerous but works.
	*
	* The whole point of this state information is to distinguish between [..arg actions..] and [char sets].
	* Char sets can only occur in lexical rules and arg actions cannot occur.
	 */
	currentRuleType int
}

// NewLexerAdaptor creates a new lexer adaptor
func NewLexerAdaptor(input antlr.CharStream) *LexerAdaptor {
	lexer := &LexerAdaptor{
		BaseLexer:         antlr.NewBaseLexer(input),
		PREQUEL_CONSTRUCT: -10,
		OPTIONS_CONSTRUCT: -11,
		currentRuleType:   antlr.TokenInvalidType,
	}
	return lexer
}

// Reset resets the lexer state
func (l *LexerAdaptor) Reset() {
	l.currentRuleType = antlr.TokenInvalidType
	l.BaseLexer.Reset()
}

func (l *LexerAdaptor) Emit() antlr.Token {
	if (l.GetType() == ANTLRv4LexerOPTIONS ||
		l.GetType() == ANTLRv4LexerTOKENS ||
		l.GetType() == ANTLRv4LexerCHANNELS) && l.currentRuleType == antlr.TokenInvalidType {
		// Enter prequel construct ending with a RBRACE.
		l.currentRuleType = l.PREQUEL_CONSTRUCT
	} else if l.GetType() == ANTLRv4LexerOPTIONS && l.currentRuleType == ANTLRv4LexerTOKEN_REF {
		l.currentRuleType = l.OPTIONS_CONSTRUCT
	} else if l.GetType() == ANTLRv4LexerRBRACE && l.currentRuleType == l.PREQUEL_CONSTRUCT {
		// Exit prequel construct.
		l.currentRuleType = antlr.TokenInvalidType
	} else if l.GetType() == ANTLRv4LexerRBRACE && l.currentRuleType == l.OPTIONS_CONSTRUCT {
		// Exit options.
		l.currentRuleType = ANTLRv4LexerTOKEN_REF
	} else if l.GetType() == ANTLRv4LexerAT && l.currentRuleType == antlr.TokenInvalidType {
		// Enter action.
		l.currentRuleType = ANTLRv4LexerAT
	} else if l.GetType() == ANTLRv4LexerSEMI && l.currentRuleType == l.OPTIONS_CONSTRUCT {
		// ';' in options { .... }. Don't change anything.
	} else if l.GetType() == ANTLRv4LexerACTION && l.currentRuleType == ANTLRv4LexerAT {
		// Exit action
		l.currentRuleType = antlr.TokenInvalidType
	} else if l.GetType() == ANTLRv4LexerID {
		firstChar := l.GetInputStream().GetText(l.TokenStartCharIndex, l.TokenStartCharIndex)
		c := len(firstChar) > 0 && firstChar[0] >= 'A' && firstChar[0] <= 'Z'
		if c {
			l.SetType(ANTLRv4LexerTOKEN_REF)
		} else {
			l.SetType(ANTLRv4LexerRULE_REF)
		}

		// If outside of the rule def.
		if l.currentRuleType == antlr.TokenInvalidType {
			// Set to inside lexer or parser rule.
			l.currentRuleType = l.GetType()
		}
	} else if l.GetType() == ANTLRv4LexerSEMI {
		// Exit rule def.
		l.currentRuleType = antlr.TokenInvalidType
	}

	return l.BaseLexer.Emit()
}

// HandleBeginArgument handles the beginning of argument constructs like [...]
// This distinguishes between character sets [abc] and semantic actions [action_code]
func (l *LexerAdaptor) HandleBeginArgument() {
	// ANTLRv4LexerTOKEN_REF context means we're in a lexer rule, so [abc] is a character set
	if l.currentRuleType == ANTLRv4LexerTOKEN_REF {
		l.PushMode(ANTLRv4LexerLexerCharSet)
		l.More()
	} else {
		// Parser rule context, so [action] is a semantic action
		l.PushMode(ANTLRv4LexerArgument)
	}
}

// HandleEndArgument handles the end of argument constructs
func (l *LexerAdaptor) HandleEndArgument() {
	l.PopMode()

	// Set token type if we're still in a mode stack
	if l.BaseLexer.ModeStackLength() > 0 {
		l.SetType(ANTLRv4LexerARGUMENT_CONTENT)
	}
}

// NextToken override to call our custom Emit logic
func (l *LexerAdaptor) NextToken() antlr.Token {
	// Get the next token using the base implementation
	token := l.BaseLexer.NextToken()
	
	// Now apply our custom emit logic to modify the token type if needed
	if token.GetTokenType() != antlr.TokenEOF {
		token = l.applyEmitLogic(token)
	}
	
	return token
}

// applyEmitLogic applies the same logic as the Java emit() method
func (l *LexerAdaptor) applyEmitLogic(token antlr.Token) antlr.Token {
	tokenType := token.GetTokenType()
	
	if (tokenType == ANTLRv4LexerOPTIONS ||
		tokenType == ANTLRv4LexerTOKENS ||
		tokenType == ANTLRv4LexerCHANNELS) && l.currentRuleType == antlr.TokenInvalidType {
		// Enter prequel construct ending with a RBRACE.
		l.currentRuleType = l.PREQUEL_CONSTRUCT
	} else if tokenType == ANTLRv4LexerOPTIONS && l.currentRuleType == ANTLRv4LexerTOKEN_REF {
		l.currentRuleType = l.OPTIONS_CONSTRUCT
	} else if tokenType == ANTLRv4LexerRBRACE && l.currentRuleType == l.PREQUEL_CONSTRUCT {
		// Exit prequel construct.
		l.currentRuleType = antlr.TokenInvalidType
	} else if tokenType == ANTLRv4LexerRBRACE && l.currentRuleType == l.OPTIONS_CONSTRUCT {
		// Exit options.
		l.currentRuleType = ANTLRv4LexerTOKEN_REF
	} else if tokenType == ANTLRv4LexerAT && l.currentRuleType == antlr.TokenInvalidType {
		// Enter action.
		l.currentRuleType = ANTLRv4LexerAT
	} else if tokenType == ANTLRv4LexerSEMI && l.currentRuleType == l.OPTIONS_CONSTRUCT {
		// ';' in options { .... }. Don't change anything.
	} else if tokenType == ANTLRv4LexerACTION && l.currentRuleType == ANTLRv4LexerAT {
		// Exit action
		l.currentRuleType = antlr.TokenInvalidType
	} else if tokenType == ANTLRv4LexerID {
		// This is the key part - convert ID to TOKEN_REF or RULE_REF
		text := token.GetText()
		if len(text) > 0 {
			firstChar := rune(text[0])
			if firstChar >= 'A' && firstChar <= 'Z' {
				// Uppercase = token reference
				token = l.setTokenType(token, ANTLRv4LexerTOKEN_REF)
			} else {
				// Lowercase = rule reference
				token = l.setTokenType(token, ANTLRv4LexerRULE_REF)
			}
		}

		// If outside of the rule def.
		if l.currentRuleType == antlr.TokenInvalidType {
			// Set to inside lexer or parser rule.
			l.currentRuleType = token.GetTokenType() // Use the new token type
		}
	} else if tokenType == ANTLRv4LexerSEMI {
		// Exit rule def.
		l.currentRuleType = antlr.TokenInvalidType
	}
	
	return token
}

// TokenTypeWrapper wraps a token and overrides its type
type TokenTypeWrapper struct {
	antlr.Token
	overriddenType int
}

func (w *TokenTypeWrapper) GetTokenType() int {
	return w.overriddenType
}

// setTokenType creates a wrapper token with the new type
func (l *LexerAdaptor) setTokenType(token antlr.Token, newType int) antlr.Token {
	return &TokenTypeWrapper{
		Token:          token,
		overriddenType: newType,
	}
}

// Helper methods for mode management
func (l *LexerAdaptor) PushMode(mode int) {
	l.BaseLexer.PushMode(mode)
}

func (l *LexerAdaptor) PopMode() int {
	return l.BaseLexer.PopMode()
}
