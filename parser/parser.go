package parser

import (
	"fmt"
)

// Parser parses tokens into an AST
type Parser struct {
	tokens   []Token
	current  int
	indentLevel int
}

// Parse converts a klo source string into an AST
func Parse(source string) (*Program, error) {
	lexer := NewLexer(source)
	tokens, err := lexer.Tokenize()
	if err != nil {
		return nil, err
	}
	
	parser := &Parser{
		tokens:  tokens,
		current: 0,
	}
	
	return parser.parseProgram()
}

func (p *Parser) parseProgram() (*Program, error) {
	program := &Program{
		Statements: []Statement{},
	}
	
	for !p.isAtEnd() {
		// Skip newlines at the beginning
		if p.check(NEWLINE) {
			p.advance()
			continue
		}
		
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		
		// Skip trailing newlines
		for p.check(NEWLINE) {
			p.advance()
		}
	}
	
	return program, nil
}

func (p *Parser) parseStatement() (Statement, error) {
	if p.check(PRINT) {
		return p.parsePrintStatement()
	}
	
	if p.check(IF) {
		return p.parseIfStatement()
	}
	
	if p.check(FOR) {
		return p.parseForStatement()
	}
	
	// Check for assignment
	if p.check(IDENTIFIER) && p.checkNext(ASSIGN) {
		return p.parseAssignmentStatement()
	}
	
	// Expression statement
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	
	return &ExpressionStatement{Expression: expr}, nil
}

func (p *Parser) parsePrintStatement() (*PrintStatement, error) {
	p.consume(PRINT, "Expected 'print'")
	
	args := []Expression{}
	
	// Parse first argument
	if !p.check(NEWLINE) && !p.isAtEnd() {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		args = append(args, expr)
		
		// Parse additional arguments separated by commas
		for p.match(COMMA) {
			expr, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			args = append(args, expr)
		}
	}
	
	return &PrintStatement{Arguments: args}, nil
}

func (p *Parser) parseAssignmentStatement() (*AssignmentStatement, error) {
	if !p.check(IDENTIFIER) {
		return nil, fmt.Errorf("expected identifier")
	}
	
	name := p.peek().Value
	p.advance() // consume identifier
	p.consume(ASSIGN, "Expected '='")
	
	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	
	return &AssignmentStatement{
		Name:  name,
		Value: value,
	}, nil
}

func (p *Parser) parseIfStatement() (*IfStatement, error) {
	p.consume(IF, "Expected 'if'")
	
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	
	p.consume(COLON, "Expected ':' after if condition")
	
	// Skip newline after colon
	if p.check(NEWLINE) {
		p.advance()
	}
	
	// Parse indented body
	body, err := p.parseIndentedBlock()
	if err != nil {
		return nil, err
	}
	
	var elseBody []Statement
	
	// Check for else clause
	if p.check(ELSE) {
		p.advance()
		p.consume(COLON, "Expected ':' after else")
		
		if p.check(NEWLINE) {
			p.advance()
		}
		
		elseBody, err = p.parseIndentedBlock()
		if err != nil {
			return nil, err
		}
	}
	
	return &IfStatement{
		Condition: condition,
		Body:      body,
		Else:      elseBody,
	}, nil
}

func (p *Parser) parseForStatement() (*ForStatement, error) {
	p.consume(FOR, "Expected 'for'")
	
	// Parse variable name (e.g., "i" in "for i in range(5)")
	if !p.check(IDENTIFIER) {
		return nil, fmt.Errorf("expected variable name after 'for'")
	}
	variable := p.peek().Value
	p.advance()
	
	p.consume(IN, "Expected 'in' after for variable")
	
	// Parse iterable (e.g., range(5))
	iterable, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	
	p.consume(COLON, "Expected ':' after for expression")
	
	// Skip newline after colon
	if p.check(NEWLINE) {
		p.advance()
	}
	
	// Parse indented body
	body, err := p.parseIndentedBlock()
	if err != nil {
		return nil, err
	}
	
	return &ForStatement{
		Variable: variable,
		Iterable: iterable,
		Body:     body,
	}, nil
}

func (p *Parser) parseIndentedBlock() ([]Statement, error) {
	statements := []Statement{}
	
	// Skip initial whitespace/indentation
	for p.check(NEWLINE) {
		p.advance()
	}
	
	// Parse statements until we hit a non-indented line or EOF
	for !p.isAtEnd() {
		// Skip empty lines
		if p.check(NEWLINE) {
			p.advance()
			continue
		}
		
		// If we hit a keyword that starts a new block (if, for, etc), we're done
		if p.check(IF) || p.check(FOR) || p.check(ELSE) {
			break
		}
		
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		
		if stmt != nil {
			statements = append(statements, stmt)
		}
		
		// Skip newlines after statement
		for p.check(NEWLINE) {
			p.advance()
		}
		
		// Break after first statement for now (simple approach)
		break
	}
	
	return statements, nil
}

func (p *Parser) parseExpression() (Expression, error) {
	return p.parseComparison()
}

func (p *Parser) parseComparison() (Expression, error) {
	expr, err := p.parseAddition()
	if err != nil {
		return nil, err
	}
	
	for p.match(GREATER, GREATER_EQ, LESS, LESS_EQ, EQUAL, NOT_EQUAL) {
		operator := p.previous().Value
		right, err := p.parseAddition()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExpression{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	
	return expr, nil
}

func (p *Parser) parseAddition() (Expression, error) {
	expr, err := p.parseMultiplication()
	if err != nil {
		return nil, err
	}
	
	for p.match(MINUS, PLUS) {
		operator := p.previous().Value
		right, err := p.parseMultiplication()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExpression{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	
	return expr, nil
}

func (p *Parser) parseMultiplication() (Expression, error) {
	expr, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}
	
	for p.match(DIVIDE, MULTIPLY, MODULO) {
		operator := p.previous().Value
		right, err := p.parsePrimary()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExpression{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	
	return expr, nil
}

func (p *Parser) parsePrimary() (Expression, error) {
	if p.match(NUMBER) {
		return &NumberLiteral{Value: p.previous().Value}, nil
	}
	
	if p.match(STRING) {
		return &StringLiteral{Value: p.previous().Value}, nil
	}
	
	if p.match(IDENTIFIER) {
		name := p.previous().Value
		
		// Check for function call like range(5)
		if p.check(LPAREN) && name == "range" {
			p.advance() // consume '('
			
			expr, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			
			p.consume(RPAREN, "Expected ')' after range argument")
			
			return &RangeExpression{End: expr}, nil
		}
		
		return &Identifier{Value: name}, nil
	}
	
	if p.match(LPAREN) {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		p.consume(RPAREN, "Expected ')' after expression")
		return expr, nil
	}
	
	return nil, fmt.Errorf("unexpected token: %v", p.peek())
}

// Helper methods
func (p *Parser) match(types ...TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) checkNext(tokenType TokenType) bool {
	if p.current + 1 >= len(p.tokens) {
		return false
	}
	return p.tokens[p.current + 1].Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) consume(tokenType TokenType, message string) error {
	if p.check(tokenType) {
		p.advance()
		return nil
	}
	
	return fmt.Errorf("%s, got %v", message, p.peek())
}
