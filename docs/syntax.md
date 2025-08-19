# klo Language Syntax Reference

## Overview

klo uses a clean, Python-inspired syntax that transpiles to efficient Go code. This document describes the complete syntax of the klo programming language.

## Comments

```klo
# This is a single-line comment
print "Hello" # Comment at end of line
```

## Variables

Variables are declared using simple assignment:

```klo
x = 42
name = "klo"
pi = 3.14159
```

Variables are dynamically assigned but statically typed at transpile time.

## Data Types

### Numbers
```klo
integer = 42
float = 3.14
negative = -10
```

### Strings
```klo
single_quotes = 'Hello'
double_quotes = "World"
```

### Identifiers
Valid variable names:
- Must start with a letter or underscore
- Can contain letters, numbers, and underscores
- Case-sensitive

```klo
valid_name = 1
_private = 2
camelCase = 3
PascalCase = 4
```

## Operators

### Arithmetic
```klo
addition = a + b
subtraction = a - b
multiplication = a * b
division = a / b
modulo = a % b
```

### Comparison
```klo
equal = a == b
not_equal = a != b
less_than = a < b
less_equal = a <= b
greater_than = a > b
greater_equal = a >= b
```

### String Concatenation
```klo
greeting = "Hello" + " " + "World"
```

## Print Statement

The `print` statement outputs values to the console:

```klo
print "Hello, World!"
print 42
print x, y, z  # Multiple values
```

## Conditional Statements

### If Statement
```klo
if x > 10:
  print "x is big"
```

### If-Else Statement
```klo
if x > 10:
  print "x is big"
else:
  print "x is small"
```

### Complex Conditions
```klo
if age >= 18:
  print "Adult"
else:
  print "Minor"

if score >= 90:
  print "A grade"
else:
  if score >= 80:
    print "B grade"
  else:
    print "Need improvement"
```

## Indentation

klo uses indentation to define code blocks, similar to Python:

- Use consistent indentation (spaces or tabs, but not mixed)
- All statements in a block must be indented at the same level
- The standard is 2 spaces per indentation level

```klo
if true:
  print "This is indented"
  if true:
    print "This is nested indentation"
  print "Back to first level"
print "This is not indented"
```

## Expressions

### Parentheses for Grouping
```klo
result = (a + b) * c
complex = ((x + y) * z) / (a - b)
```

### Operator Precedence
1. Parentheses `()`
2. Multiplication, Division, Modulo `*`, `/`, `%`
3. Addition, Subtraction `+`, `-`
4. Comparison `<`, `<=`, `>`, `>=`
5. Equality `==`, `!=`

```klo
# This is evaluated as: ((2 + 3) * 4) > (10 / 2)
result = 2 + 3 * 4 > 10 / 2
```

## Future Features (Planned)

### Functions
```klo
def greet(name):
  print "Hello, " + name
  return "Greeting sent"

result = greet("Alice")
```

### Loops
```klo
# For loops
for i in range(5):
  print i

for item in items:
  print item

# While loops
while x < 10:
  x = x + 1
  print x
```

### Arrays/Lists
```klo
numbers = [1, 2, 3, 4, 5]
names = ["Alice", "Bob", "Charlie"]

first = numbers[0]
length = len(numbers)
```

### Objects/Maps
```klo
person = {
  "name": "Alice",
  "age": 30,
  "city": "New York"
}

name = person["name"]
```

## Error Handling

Current error handling is limited to compile-time errors. The transpiler will catch:

- Syntax errors
- Undefined tokens
- Malformed expressions

## Best Practices

1. **Use meaningful variable names**
   ```klo
   # Good
   user_age = 25
   total_price = 99.99
   
   # Avoid
   x = 25
   tp = 99.99
   ```

2. **Keep expressions simple**
   ```klo
   # Good
   subtotal = price * quantity
   tax = subtotal * tax_rate
   total = subtotal + tax
   
   # Avoid
   total = (price * quantity) + ((price * quantity) * tax_rate)
   ```

3. **Use consistent indentation**
   ```klo
   # Good
   if condition:
     do_something()
     if nested_condition:
       do_nested_thing()
   
   # Avoid mixing indentation styles
   ```

4. **Add comments for complex logic**
   ```klo
   # Calculate compound interest
   amount = principal * (1 + rate) ** time
   ```

## Examples

### Hello World
```klo
print "Hello, klo!"
```

### Calculator
```klo
a = 10
b = 5

sum = a + b
difference = a - b
product = a * b
quotient = a / b

print "Sum:", sum
print "Difference:", difference
print "Product:", product
print "Quotient:", quotient
```

### Conditional Logic
```klo
score = 85

if score >= 90:
  grade = "A"
else:
  if score >= 80:
    grade = "B"
  else:
    if score >= 70:
      grade = "C"
    else:
      grade = "F"

print "Grade:", grade
```
