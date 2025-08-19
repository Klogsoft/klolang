# Installation and Usage Guide

## Quick Start

### 1. Install Go
Download and install Go from https://golang.org/dl/

### 2. Install klo
```bash
# Clone the repository
git clone https://github.com/mesut/klo.git
cd klo

# Build the binary
go build -o klo

# Add to PATH (Linux/macOS)
sudo mv klo /usr/local/bin/

# Add to PATH (Windows)
mkdir C:\tools\bin
copy klo.exe C:\tools\bin\klo.exe
[Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\tools\bin", "User")
```

### 3. Run your first klo program

Create a file called `hello.klo`:
```klo
print "Hello, klo!"
x = 5 + 3
print "x =", x
```

Run it:
```bash
klo hello.klo
```

## Command Line Usage

```bash
# Run a klo file
klo script.klo

# Transpile only (don't execute)
klo --transpile script.klo

# Save transpiled Go code to file
klo --transpile --output script.go script.klo

# Verbose output
klo --verbose script.klo

# Show version
klo version

# Show help
klo --help
```

## Language Examples

### Variables and Arithmetic
```klo
name = "Alice"
age = 25
next_year = age + 1

print "Name:", name
print "Age:", age
print "Next year:", next_year
```

### Conditionals
```klo
score = 85

if score >= 90:
  print "Excellent!"
else:
  if score >= 70:
    print "Good!"
  else:
    print "Needs improvement"
```

### Complex Expressions
```klo
a = 10
b = 5
c = 2

result = (a + b) * c
remainder = a % b
is_greater = result > 20

print "Result:", result
print "Remainder:", remainder
print "Is greater than 20:", is_greater
```

## Integration with Go

klo transparently converts your code to Go. You can see the generated Go code:

```bash
klo --transpile --output output.go your_script.klo
cat output.go
```

This allows you to:
- Learn Go gradually by seeing the transpiled output
- Debug issues by examining the generated Go code
- Understand the performance characteristics of your klo code

## Editor Support

### VS Code
While dedicated klo extensions are planned, you can use:
1. Generic syntax highlighting for similar languages
2. Go extension for inspecting transpiled code

### Vim/Neovim
Basic syntax highlighting can be added by treating `.klo` files as Python-like syntax.

## Performance

klo programs have the same performance as equivalent Go programs because:
- klo transpiles directly to Go source code
- No runtime overhead or interpretation
- Full access to Go's optimizing compiler
- Native binary output

## Debugging

To debug klo programs:
1. Use `--transpile` to see the generated Go code
2. Run the Go code directly with `go run`
3. Use Go debugging tools like Delve on the transpiled code

## Learning Path

1. **Start with basics**: variables, print statements
2. **Add conditionals**: if/else statements
3. **Use arithmetic**: mathematical expressions
4. **Explore transpilation**: see how klo becomes Go
5. **Contribute**: help add new features to klo

## Next Steps

- Join our community discussions
- Try the examples in the `examples/` directory
- Read the [syntax documentation](docs/syntax.md)
- Contribute new features or improvements
