package parser

// AST Node types
type Node interface {
	String() string
}

// Program represents the root of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	return "Program"
}

// Statement interface
type Statement interface {
	Node
	statementNode()
}

// Expression interface
type Expression interface {
	Node
	expressionNode()
}

// PrintStatement represents a print statement
type PrintStatement struct {
	Arguments []Expression
}

func (ps *PrintStatement) statementNode() {}
func (ps *PrintStatement) String() string { return "PrintStatement" }

// AssignmentStatement represents variable assignment
type AssignmentStatement struct {
	Name  string
	Value Expression
}

func (as *AssignmentStatement) statementNode() {}
func (as *AssignmentStatement) String() string { return "AssignmentStatement" }

// IfStatement represents conditional statement
type IfStatement struct {
	Condition Expression
	Body      []Statement
	Else      []Statement
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) String() string { return "IfStatement" }

// ForStatement represents for loop statement
type ForStatement struct {
	Variable string      // loop variable (e.g., "i")
	Iterable Expression  // what to iterate over (e.g., range(5))
	Body     []Statement // loop body
}

func (fs *ForStatement) statementNode() {}
func (fs *ForStatement) String() string { return "ForStatement" }

// RangeExpression represents range(n) function
type RangeExpression struct {
	End Expression // end value
}

func (re *RangeExpression) expressionNode() {}
func (re *RangeExpression) String() string { return "range(...)" }

// ExpressionStatement wraps expressions used as statements
type ExpressionStatement struct {
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string { return "ExpressionStatement" }

// Identifier represents variable names
type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) String() string { return i.Value }

// StringLiteral represents string values
type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) String() string { return "\"" + sl.Value + "\"" }

// NumberLiteral represents numeric values
type NumberLiteral struct {
	Value string
}

func (nl *NumberLiteral) expressionNode() {}
func (nl *NumberLiteral) String() string { return nl.Value }

// BinaryExpression represents binary operations
type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) expressionNode() {}
func (be *BinaryExpression) String() string {
	return "(" + be.Left.String() + " " + be.Operator + " " + be.Right.String() + ")"
}
