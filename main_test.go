package main

import (
	"testing"
	"github.com/singleservingfriend/klo/parser"
	"github.com/singleservingfriend/klo/transpiler"
)

func TestBasicParsing(t *testing.T) {
	source := `print "Hello, World!"`
	
	program, err := parser.Parse(source)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	
	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}
	
	printStmt, ok := program.Statements[0].(*parser.PrintStatement)
	if !ok {
		t.Fatalf("Expected PrintStatement, got %T", program.Statements[0])
	}
	
	if len(printStmt.Arguments) != 1 {
		t.Fatalf("Expected 1 argument, got %d", len(printStmt.Arguments))
	}
}

func TestVariableAssignment(t *testing.T) {
	source := `x = 42`
	
	program, err := parser.Parse(source)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	
	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}
	
	assignStmt, ok := program.Statements[0].(*parser.AssignmentStatement)
	if !ok {
		t.Fatalf("Expected AssignmentStatement, got %T", program.Statements[0])
	}
	
	if assignStmt.Name != "x" {
		t.Fatalf("Expected variable name 'x', got '%s'", assignStmt.Name)
	}
}

func TestArithmetic(t *testing.T) {
	source := `result = 5 + 3 * 2`
	
	program, err := parser.Parse(source)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	
	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}
	
	assignStmt, ok := program.Statements[0].(*parser.AssignmentStatement)
	if !ok {
		t.Fatalf("Expected AssignmentStatement, got %T", program.Statements[0])
	}
	
	_, ok = assignStmt.Value.(*parser.BinaryExpression)
	if !ok {
		t.Fatalf("Expected BinaryExpression, got %T", assignStmt.Value)
	}
}

func TestForLoop(t *testing.T) {
	source := `for i in range(5):
  print i`
	
	program, err := parser.Parse(source)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	
	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}
	
	forStmt, ok := program.Statements[0].(*parser.ForStatement)
	if !ok {
		t.Fatalf("Expected ForStatement, got %T", program.Statements[0])
	}
	
	if forStmt.Variable != "i" {
		t.Fatalf("Expected variable name 'i', got '%s'", forStmt.Variable)
	}
}

func TestTranspilation(t *testing.T) {
	source := `x = 42
print x`
	
	program, err := parser.Parse(source)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	
	goCode := transpiler.GenerateGoCode(program)
	
	// Basic checks that Go code was generated
	if len(goCode) == 0 {
		t.Fatal("No Go code generated")
	}
	
	// Should contain package declaration
	if !contains(goCode, "package main") {
		t.Fatal("Generated code missing package declaration")
	}
	
	// Should contain main function
	if !contains(goCode, "func main()") {
		t.Fatal("Generated code missing main function")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr || len(s) > len(substr) && contains(s[1:], substr)
}
