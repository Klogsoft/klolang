package parser

import (
	"fmt"
)

// Token represents a single token in the klo language
type Token struct {
	Type    TokenType
	Value   string
	Line    int
	Column  int
}

// TokenType represents the type of token
type TokenType int

const (
	// Basic tokens
	EOF TokenType = iota
	NEWLINE
	INDENT
	DEDENT
	
	// Literals
	IDENTIFIER
	STRING
	NUMBER
	
	// Keywords
	PRINT
	IF
	ELSE
	DEF
	FOR
	IN
	WHILE
	RETURN
	
	// Operators
	ASSIGN     // =
	PLUS       // +
	MINUS      // -
	MULTIPLY   // *
	DIVIDE     // /
	MODULO     // %
	
	// Comparison
	EQUAL      // ==
	NOT_EQUAL  // !=
	LESS       // <
	LESS_EQ    // <=
	GREATER    // >
	GREATER_EQ // >=
	
	// Punctuation
	LPAREN     // (
	RPAREN     // )
	LBRACKET   // [
	RBRACKET   // ]
	COMMA      // ,
	COLON      // :
	DOT        // .
)

// Lexer tokenizes klo source code
type Lexer struct {
	input    string
	position int
	line     int
	column   int
	tokens   []Token
}

// NewLexer creates a new lexer instance
func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		line:   1,
		column: 1,
	}
}

// Tokenize converts the input string into tokens
func (l *Lexer) Tokenize() ([]Token, error) {
	for l.position < len(l.input) {
		if err := l.scanToken(); err != nil {
			return nil, err
		}
	}
	
	l.addToken(EOF, "")
	return l.tokens, nil
}

func (l *Lexer) scanToken() error {
	ch := l.currentChar()
	
	switch {
	case ch == '\n' || ch == '\r':
		// Handle both Unix (\n) and Windows (\r\n) line endings
		if ch == '\r' && l.peekChar() == '\n' {
			// Windows line ending, skip \r and \n
			l.advance()
			l.advance()
		} else {
			// Unix line ending, just skip \n
			l.advance()
		}
		l.addToken(NEWLINE, "\n")
		l.line++
		l.column = 1
		return nil
		
	case ch == ' ' || ch == '\t':
		l.skipWhitespace()
		return nil
		
	case ch == '#':
		l.skipComment()
		return nil
		
	case ch == '"' || ch == '\'':
		return l.scanString()
		
	case isDigit(ch):
		return l.scanNumber()
		
	case isAlpha(ch):
		return l.scanIdentifier()
		
	case ch == '=':
		if l.peekChar() == '=' {
			l.addToken(EQUAL, "==")
			l.advance()
			l.advance()
		} else {
			l.addToken(ASSIGN, "=")
			l.advance()
		}
		return nil
		
	case ch == '!':
		if l.peekChar() == '=' {
			l.addToken(NOT_EQUAL, "!=")
			l.advance()
			l.advance()
		} else {
			return fmt.Errorf("unexpected character '!' at line %d, column %d", l.line, l.column)
		}
		return nil
		
	case ch == '<':
		if l.peekChar() == '=' {
			l.addToken(LESS_EQ, "<=")
			l.advance()
			l.advance()
		} else {
			l.addToken(LESS, "<")
			l.advance()
		}
		return nil
		
	case ch == '>':
		if l.peekChar() == '=' {
			l.addToken(GREATER_EQ, ">=")
			l.advance()
			l.advance()
		} else {
			l.addToken(GREATER, ">")
			l.advance()
		}
		return nil
		
	case ch == '+':
		l.addToken(PLUS, "+")
		l.advance()
		return nil
		
	case ch == '-':
		l.addToken(MINUS, "-")
		l.advance()
		return nil
		
	case ch == '*':
		l.addToken(MULTIPLY, "*")
		l.advance()
		return nil
		
	case ch == '/':
		l.addToken(DIVIDE, "/")
		l.advance()
		return nil
		
	case ch == '%':
		l.addToken(MODULO, "%")
		l.advance()
		return nil
		
	case ch == '(':
		l.addToken(LPAREN, "(")
		l.advance()
		return nil
		
	case ch == ')':
		l.addToken(RPAREN, ")")
		l.advance()
		return nil
		
	case ch == '[':
		l.addToken(LBRACKET, "[")
		l.advance()
		return nil
		
	case ch == ']':
		l.addToken(RBRACKET, "]")
		l.advance()
		return nil
		
	case ch == ',':
		l.addToken(COMMA, ",")
		l.advance()
		return nil
		
	case ch == ':':
		l.addToken(COLON, ":")
		l.advance()
		return nil
		
	case ch == '.':
		l.addToken(DOT, ".")
		l.advance()
		return nil
		
	default:
		return fmt.Errorf("unexpected character '%c' at line %d, column %d", ch, l.line, l.column)
	}
}

func (l *Lexer) currentChar() byte {
	if l.position >= len(l.input) {
		return 0
	}
	return l.input[l.position]
}

func (l *Lexer) peekChar() byte {
	if l.position+1 >= len(l.input) {
		return 0
	}
	return l.input[l.position+1]
}

func (l *Lexer) advance() {
	l.position++
	l.column++
}

func (l *Lexer) addToken(tokenType TokenType, value string) {
	l.tokens = append(l.tokens, Token{
		Type:   tokenType,
		Value:  value,
		Line:   l.line,
		Column: l.column,
	})
}

func (l *Lexer) skipWhitespace() {
	for l.position < len(l.input) && (l.currentChar() == ' ' || l.currentChar() == '\t') {
		l.advance()
	}
}

func (l *Lexer) skipComment() {
	for l.position < len(l.input) && l.currentChar() != '\n' {
		l.advance()
	}
}

func (l *Lexer) scanString() error {
	quote := l.currentChar()
	l.advance() // Skip opening quote
	
	start := l.position
	for l.position < len(l.input) && l.currentChar() != quote {
		if l.currentChar() == '\n' {
			l.line++
			l.column = 1
		}
		l.advance()
	}
	
	if l.position >= len(l.input) {
		return fmt.Errorf("unterminated string at line %d", l.line)
	}
	
	value := l.input[start:l.position]
	l.advance() // Skip closing quote
	
	l.addToken(STRING, value)
	return nil
}

func (l *Lexer) scanNumber() error {
	start := l.position
	
	for l.position < len(l.input) && (isDigit(l.currentChar()) || l.currentChar() == '.') {
		l.advance()
	}
	
	value := l.input[start:l.position]
	l.addToken(NUMBER, value)
	return nil
}

func (l *Lexer) scanIdentifier() error {
	start := l.position
	
	for l.position < len(l.input) && (isAlnum(l.currentChar()) || l.currentChar() == '_') {
		l.advance()
	}
	
	value := l.input[start:l.position]
	tokenType := identifierType(value)
	
	l.addToken(tokenType, value)
	return nil
}

func identifierType(value string) TokenType {
	keywords := map[string]TokenType{
		"print":  PRINT,
		"if":     IF,
		"else":   ELSE,
		"def":    DEF,
		"for":    FOR,
		"in":     IN,
		"while":  WHILE,
		"return": RETURN,
		"range":  IDENTIFIER, // range is treated as identifier/function
	}
	
	if tokenType, exists := keywords[value]; exists {
		return tokenType
	}
	
	return IDENTIFIER
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isAlnum(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}
