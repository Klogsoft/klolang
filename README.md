# 🚀 klo Programming Language

<div align="center">

![klo Logo](https://img.shields.io/badge/klo-v0.1.0-blue)
![Go Version](https://img.shields.io/badge/go-1.21+-green)
![License](https://img.shields.io/badge/license-MIT-green)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)

**A minimalist programming language that makes coding simple and fun**

[🎯 Quick Start](#-quick-start) • [📖 Documentation](#-documentation) • [💡 Examples](#-examples) • [🤝 Contributing](#-contributing)

</div>

---

## 🌟 What is klo?

**klo** is a modern, minimalist programming language designed to be:
- **Easy to learn** - Python-like syntax that's beginner-friendly
- **Fast to execute** - Transpiles to optimized Go code
- **Simple to use** - One command to run your programs
- **Powerful** - Full access to Go's ecosystem and performance

### 🎯 Perfect for:
- � **Learning programming** - Clear, readable syntax
- 🔬 **Prototyping** - Quick idea testing
- � **Scripting** - Automation and tools
- � **Teaching** - Educational environments

---

## ⚡ Quick Start

### 1️⃣ Install Go
```bash
# Download from: https://golang.org/dl/
# Or use package manager:

# Windows (Chocolatey)
choco install golang

# macOS (Homebrew)
brew install go

# Ubuntu/Debian
sudo apt install golang-go
```

### 2️⃣ Install klo
```bash
# Clone the repository
git clone https://github.com/singleservingfriend/klo.git
cd klo

# Build klo
go build -o klo

# Add to PATH (choose your platform)
```

**Windows:**
```powershell
# Create tools directory
mkdir C:\tools\bin
copy klo.exe C:\tools\bin\

# Add to PATH permanently
[Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\tools\bin", "User")

# Restart terminal and verify
klo version
```

**Linux/macOS:**
```bash
# Add to system PATH
sudo mv klo /usr/local/bin/

# Verify installation
klo version
```

### 3️⃣ Write Your First Program
Create `hello.klo`:
```klo
print "Hello, World!"
name = "klo"
print "Welcome to", name
```

### 4️⃣ Run It!
```bash
klo hello.klo
```

**Output:**
```
Hello, World!
Welcome to klo
```

🎉 **Congratulations! You just ran your first klo program!**

---

## 🛠️ Current Language Features

### ✅ What klo Can Do Right Now

#### 📝 **Variables & Basic Operations**
```klo
# Variables
name = "Alice"
age = 25
height = 5.8

# Arithmetic
result = (10 + 5) * 2
remainder = 17 % 3
```

#### 🖨️ **Print Statements**
```klo
print "Simple text"
print "Age:", age
print name, age, height  # Multiple values
```

#### ⚖️ **Conditionals**
```klo
score = 85

if score >= 90:
  print "Excellent!"
else:
  if score >= 70:
    print "Good job!"
  else:
    print "Keep practicing!"
```

#### 🔄 **For Loops**
```klo
# Count from 0 to 4
for i in range(5):
  print "Count:", i

# Practical example
for num in range(1, 6):
  square = num * num
  print num, "squared is", square
```

#### 🔢 **Mathematical Operations**
```klo
# All basic operators supported
a = 10
b = 3

print "Addition:", a + b      # 13
print "Subtraction:", a - b   # 7  
print "Multiplication:", a * b # 30
print "Division:", a / b      # 3
print "Modulo:", a % b        # 1

# Comparisons
print a > b    # true
print a == b   # false
print a <= b   # false
```

#### 💬 **Comments**
```klo
# This is a single-line comment
x = 5  # Comments can also be at the end of lines

# Use comments to explain your code
# They are ignored during execution
```

### 🚧 **Coming Soon** (Help us build these!)
- **Functions**: `def myFunction(param):`
- **While loops**: `while condition:`
- **Lists/Arrays**: `items = [1, 2, 3]`
- **String operations**: Advanced text manipulation
- **File I/O**: Reading and writing files
- **Error handling**: Try/catch mechanisms

---

## 📖 Documentation

### 🎮 Command Line Usage
```bash
# Run a klo program
klo script.klo

# See the generated Go code (great for learning!)
klo --transpile --output output.go script.klo

# Run with detailed output
klo --verbose script.klo

# Get help
klo --help

# Check version
klo version
```

### 🎯 Language Syntax Guide

#### Variables
```klo
# Numbers
age = 25
price = 19.99

# Strings  
name = "John"
message = 'Hello World'

# Expressions
total = price * 2
full_name = "Mr. " + name
```

#### Control Flow
```klo
# If statements
if age >= 18:
  print "Adult"

# If-else chains
if temperature > 30:
  print "Hot"
else:
  if temperature > 20:
    print "Warm"
  else:
    print "Cold"
```

#### Loops
```klo
# Simple counting
for i in range(3):
  print "Iteration", i

# Practical example: multiplication table
number = 7
for i in range(1, 11):
  result = number * i
  print number, "x", i, "=", result
```

---

## 💡 Examples

### 🧮 Calculator
```klo
print "Simple Calculator"

a = 15
b = 4

print "Numbers:", a, "and", b
print "Sum:", a + b
print "Difference:", a - b  
print "Product:", a * b
print "Division:", a / b

if a > b:
  print a, "is greater than", b
```

### 🔢 Fibonacci Sequence
```klo
print "Fibonacci sequence:"

a = 0
b = 1

print a
print b

# Generate next 8 numbers
for i in range(8):
  c = a + b
  print c
  a = b
  b = c
```

### 🎲 Number Guessing Game Concept
```klo
print "Number guessing game"

secret = 42
guess = 35

print "Secret number is:", secret
print "Your guess is:", guess

if guess == secret:
  print "Congratulations! You guessed it!"
else:
  if guess < secret:
    print "Too low! Try higher."
  else:
    print "Too high! Try lower."
```

---

## 🚀 How klo Works

klo uses **transpilation** - it converts your klo code into Go code, then compiles and runs it:

```klo
# Your klo code
for i in range(3):
  print "Hello", i
```

**Becomes this Go code:**
```go
package main
import "fmt"
func main() {
    for i := 0; i < 3; i++ {
        fmt.Println("Hello", i)
    }
}
```

**Benefits:**
- 🏃‍♂️ **Go's speed** - Your programs run as fast as native Go
- 📚 **Go's ecosystem** - Access to Go's standard library
- 🛡️ **Go's safety** - Memory safe and type safe
- 🎯 **Simple syntax** - Easier to write and read

---

## 🤝 Contributing & Community

### 🌍 **Open Source Development**
klo is **100% open source** and developed by the community! We believe in:
- **Collaborative development** - Everyone can contribute
- **Transparency** - All development happens in the open
- **Learning together** - Share knowledge and grow

### 🎯 **How You Can Help**

#### 🐛 **Report Issues**
Found a bug? Have a suggestion? [Open an issue!](https://github.com/singleservingfriend/klo/issues)

#### 💻 **Contribute Code**
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes
4. Add tests if needed
5. Commit: `git commit -m 'Add amazing feature'`
6. Push: `git push origin feature/amazing-feature`
7. Open a Pull Request

#### 📝 **Improve Documentation**
- Fix typos
- Add examples
- Improve explanations
- Translate to other languages

#### 🎨 **Create Examples**
Share cool programs you've written in klo!

### 🎓 **Learning Resources**
- **Examples**: Check the [`examples/`](examples/) directory
- **Syntax Guide**: Read [`docs/syntax.md`](docs/syntax.md)
- **Contributing**: See [`CONTRIBUTING.md`](CONTRIBUTING.md)

---

## 📁 Project Structure

```
klo/
├── 📄 main.go              # CLI entry point
├── 📁 parser/              # Language parser
│   ├── lexer.go           # Tokenizer
│   ├── parser.go          # Syntax analyzer  
│   └── ast.go             # Abstract syntax tree
├── 📁 transpiler/          # klo → Go converter
│   └── generator.go       # Code generation
├── 📁 examples/            # Example programs
│   ├── hello.klo         # Basic examples
│   ├── calculator.klo    # Math operations
│   └── fibonacci.klo     # Algorithms
├── 📁 docs/               # Documentation
└── 📄 README.md           # This file
```

---

## 🎯 Roadmap

### ✅ **Completed** (v0.1.0)
- [x] Basic syntax (variables, print, math)
- [x] Conditional statements (if/else)
- [x] For loops with range()
- [x] CLI tool with transpilation
- [x] Cross-platform support
- [x] Comprehensive documentation

### 🚧 **In Progress** (v0.2.0)
- [ ] Functions and parameters
- [ ] While loops
- [ ] Better error messages
- [ ] More built-in functions

### 🔮 **Future** (v0.3.0+)
- [ ] Arrays and data structures
- [ ] String manipulation
- [ ] File I/O operations
- [ ] Package system
- [ ] VS Code extension
- [ ] Online playground
- [ ] Interactive REPL

---

## 🆘 Getting Help

### 💬 **Community Support**
- **GitHub Issues**: [Report bugs or ask questions](https://github.com/singleservingfriend/klo/issues)
- **Discussions**: [Community discussions](https://github.com/singleservingfriend/klo/discussions)

### 📚 **Documentation**
- **Syntax Reference**: [`docs/syntax.md`](docs/syntax.md)
- **Installation Guide**: [`SETUP.md`](SETUP.md)
- **Usage Examples**: [`USAGE.md`](USAGE.md)

### 🐛 **Common Issues**
1. **"klo command not found"** → Check if klo is in your PATH
2. **"go command not found"** → Install Go first
3. **Parse errors** → Check syntax, especially indentation

---

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

**TL;DR**: You can use, modify, and distribute klo freely!

---

## 👥 Authors & Community

**Created by**: Klogsoft  
**Maintained by**: The klo community

### 🌟 **Special Thanks**
- All contributors who help improve klo
- The Go team for the amazing Go language
- Everyone who uses and shares klo

---

## 🔗 Links

- 🌐 **Website and Documentation**: https://klogsoft.com/klo

---

<div align="center">

**Made with ❤️ by the klo community**

*"Making programming simple, one line at a time"*

⭐ **Star this repo if you like klo!** ⭐

</div>
