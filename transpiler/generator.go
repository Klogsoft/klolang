package transpiler

import (
	"fmt"
	"strings"
	"github.com/singleservingfriend/klo/parser"
)

// GenerateGoCode converts a klo AST to Go source code
func GenerateGoCode(program *parser.Program) string {
	generator := &GoGenerator{
		indent: 0,
	}
	
	return generator.generateProgram(program)
}

// GoGenerator handles the conversion from AST to Go code
type GoGenerator struct {
	indent int
}

func (g *GoGenerator) generateProgram(program *parser.Program) string {
	var output strings.Builder
	
	// Add package declaration and imports
	output.WriteString("package main\n\n")
	output.WriteString("import (\n")
	output.WriteString("\t\"fmt\"\n")
	output.WriteString(")\n\n")
	
	// Add main function
	output.WriteString("func main() {\n")
	g.indent++
	
	// Generate statements
	for _, stmt := range program.Statements {
		code := g.generateStatement(stmt)
		if code != "" {
			output.WriteString(g.indentString() + code + "\n")
		}
	}
	
	g.indent--
	output.WriteString("}\n")
	
	return output.String()
}

func (g *GoGenerator) generateStatement(stmt parser.Statement) string {
	switch s := stmt.(type) {
	case *parser.PrintStatement:
		return g.generatePrintStatement(s)
	case *parser.AssignmentStatement:
		return g.generateAssignmentStatement(s)
	case *parser.IfStatement:
		return g.generateIfStatement(s)
	case *parser.ForStatement:
		return g.generateForStatement(s)
	case *parser.ExpressionStatement:
		return g.generateExpression(s.Expression)
	default:
		return fmt.Sprintf("// Unknown statement: %T", stmt)
	}
}

func (g *GoGenerator) generatePrintStatement(stmt *parser.PrintStatement) string {
	if len(stmt.Arguments) == 0 {
		return "fmt.Println()"
	}
	
	args := make([]string, len(stmt.Arguments))
	for i, arg := range stmt.Arguments {
		args[i] = g.generateExpression(arg)
	}
	
	if len(args) == 1 {
		return fmt.Sprintf("fmt.Println(%s)", args[0])
	}
	
	// Multiple arguments - use fmt.Println with all args
	return fmt.Sprintf("fmt.Println(%s)", strings.Join(args, ", "))
}

func (g *GoGenerator) generateAssignmentStatement(stmt *parser.AssignmentStatement) string {
	value := g.generateExpression(stmt.Value)
	return fmt.Sprintf("%s := %s", stmt.Name, value)
}

func (g *GoGenerator) generateIfStatement(stmt *parser.IfStatement) string {
	var output strings.Builder
	
	condition := g.generateExpression(stmt.Condition)
	output.WriteString(fmt.Sprintf("if %s {\n", condition))
	
	g.indent++
	for _, bodyStmt := range stmt.Body {
		code := g.generateStatement(bodyStmt)
		if code != "" {
			output.WriteString(g.indentString() + code + "\n")
		}
	}
	g.indent--
	
	if len(stmt.Else) > 0 {
		output.WriteString(g.indentString() + "} else {\n")
		g.indent++
		for _, elseStmt := range stmt.Else {
			code := g.generateStatement(elseStmt)
			if code != "" {
				output.WriteString(g.indentString() + code + "\n")
			}
		}
		g.indent--
	}
	
	output.WriteString(g.indentString() + "}")
	
	return output.String()
}

func (g *GoGenerator) generateForStatement(stmt *parser.ForStatement) string {
	var output strings.Builder
	
	// Check if iterable is range expression
	if rangeExpr, ok := stmt.Iterable.(*parser.RangeExpression); ok {
		end := g.generateExpression(rangeExpr.End)
		output.WriteString(fmt.Sprintf("for %s := 0; %s < %s; %s++ {\n", 
			stmt.Variable, stmt.Variable, end, stmt.Variable))
	} else {
		// For other iterables (future feature)
		output.WriteString(fmt.Sprintf("// TODO: Implement iteration over %T\n", stmt.Iterable))
		return output.String()
	}
	
	g.indent++
	for _, bodyStmt := range stmt.Body {
		code := g.generateStatement(bodyStmt)
		if code != "" {
			output.WriteString(g.indentString() + code + "\n")
		}
	}
	g.indent--
	
	output.WriteString(g.indentString() + "}")
	
	return output.String()
}

func (g *GoGenerator) generateExpression(expr parser.Expression) string {
	switch e := expr.(type) {
	case *parser.Identifier:
		return e.Value
	case *parser.StringLiteral:
		return fmt.Sprintf("\"%s\"", e.Value)
	case *parser.NumberLiteral:
		return e.Value
	case *parser.BinaryExpression:
		return g.generateBinaryExpression(e)
	case *parser.RangeExpression:
		// Range expressions are handled in for loops
		return g.generateExpression(e.End)
	default:
		return fmt.Sprintf("/* Unknown expression: %T */", expr)
	}
}

func (g *GoGenerator) generateBinaryExpression(expr *parser.BinaryExpression) string {
	left := g.generateExpression(expr.Left)
	right := g.generateExpression(expr.Right)
	
	// Handle string concatenation
	if expr.Operator == "+" {
		// If either operand is a string literal, treat as string concatenation
		if g.isStringExpression(expr.Left) || g.isStringExpression(expr.Right) {
			return fmt.Sprintf("fmt.Sprintf(\"%%v%%v\", %s, %s)", left, right)
		}
	}
	
	return fmt.Sprintf("(%s %s %s)", left, expr.Operator, right)
}

func (g *GoGenerator) isStringExpression(expr parser.Expression) bool {
	switch expr.(type) {
	case *parser.StringLiteral:
		return true
	default:
		return false
	}
}

func (g *GoGenerator) indentString() string {
	return strings.Repeat("\t", g.indent)
}
