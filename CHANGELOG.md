# Changelog

All notable changes to the klo programming language will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- Function definitions (`def`)
- Loop statements (`for`, `while`)
- Array/list support
- Object/map support
- Import system
- Standard library functions

## [0.1.0] - 2025-08-19

### Added
- Initial release of klo programming language
- Basic lexer and parser
- Transpilation to Go code
- CLI tool (`klo`) for running .klo files
- Print statements
- Variable assignments
- If-else conditional statements
- Arithmetic expressions (`+`, `-`, `*`, `/`, `%`)
- Comparison operators (`==`, `!=`, `<`, `<=`, `>`, `>=`)
- String and number literals
- Comments (single-line with `#`)
- Basic error handling and reporting
- Documentation and examples

### Features
- **Print statements**: `print "Hello, World!"`
- **Variables**: `x = 42`
- **Arithmetic**: `result = (a + b) * c`
- **Conditionals**: `if x > 5: print "big"`
- **String operations**: `name = "Alice"`
- **Comments**: `# This is a comment`

### CLI Options
- `klo file.klo` - Run a klo file
- `klo --transpile file.klo` - Transpile only, don't execute
- `klo --verbose file.klo` - Show verbose output
- `klo --output result.go file.klo` - Specify output file
- `klo --help` - Show help information
- `klo --version` - Show version information

### Examples
- `hello.klo` - Basic hello world program
- `calculator.klo` - Simple arithmetic operations
- `fibonacci.klo` - Generate Fibonacci sequence

### Technical Details
- Written in Go 1.21+
- Uses recursive descent parser
- Generates clean, readable Go code
- Supports Windows, macOS, and Linux
- MIT Licensed
