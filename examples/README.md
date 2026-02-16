# Examples Directory

This folder contains additional Go examples for learning and practice.

## Running the Examples

### Hello World - The Basics
```bash
go run examples/hello-world.go
```
**What you'll learn:**
- Basic output with `fmt.Println`
- String formatting with `Printf`

---

### Calculator - Functions & Arithmetic
```bash
go run examples/calculator.go
```
**What you'll learn:**
- Function definition and calling
- Arithmetic operations
- Loop-based calculations (power function)
- Slices and iteration

**Functions included:**
- `add(x, y)` – Addition
- `subtract(x, y)` – Subtraction
- `multiply(x, y)` – Multiplication
- `divide(x, y)` – Division with error handling
- `modulo(x, y)` – Modulo operation
- `power(base, exp)` – Exponentiation
- `average(numbers)` – Slice average calculation

---

### Greetings - String Manipulation
```bash
go run examples/greetings.go
```
**What you'll learn:**
- String concatenation
- String formatting with `Sprintf`
- String package functions (ToUpper, ToLower, Split)
- String reversal algorithms
- Functions returning strings

**Functions included:**
- `greet(name)` – Personalized greeting
- `reverseString(s)` – String reversal
- `displayGreeting(name, interest)` – Multi-line output

---

## Progressive Learning Path

1. **Start with:** `hello-world.go` – Understand basic output
2. **Then:** `calculator.go` – Learn functions and loops
3. **Next:** `greetings.go` – Master string manipulation

---

## Modify & Experiment

Try these exercises:

### Calculator.go Modifications
- Add a `sqrt()` function using Newton's method
- Add a `factorial()` function
- Modify `average()` to also calculate median

### Greetings.go Modifications
- Add a `countVowels(s)` function
- Create a `titleCase(s)` function
- Build a `generatePassword(length)` function

---

## Running All Examples

```bash
echo "=== Hello World ===" && go run examples/hello-world.go
echo -e "\n=== Calculator ===" && go run examples/calculator.go
echo -e "\n=== Greetings ===" && go run examples/greetings.go
```

---

## Building Executables

Convert any example to a standalone program:

```bash
go build -o hello examples/hello-world.go
./hello

go build -o calc examples/calculator.go
./calc

go build -o greet examples/greetings.go
./greet
```

---

Happy learning! 🚀
