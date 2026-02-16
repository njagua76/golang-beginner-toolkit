package main

import (
	"fmt"
)

func RunCalculator() {
	fmt.Println("=== Simple Calculator ===")

	// Basic arithmetic operations
	a := 20
	b := 8

	fmt.Printf("Numbers: %d and %d\n", a, b)
	fmt.Printf("Addition:       %d + %d = %d\n", a, b, calcAdd(a, b))
	fmt.Printf("Subtraction:    %d - %d = %d\n", a, b, calcSubtract(a, b))
	fmt.Printf("Multiplication: %d × %d = %d\n", a, b, calcMultiply(a, b))
	fmt.Printf("Division:       %d ÷ %d = %d\n", a, b, calcDivide(a, b))
	fmt.Printf("Modulo:         %d %% %d = %d\n", a, b, calcModulo(a, b))

	fmt.Println("\n=== Advanced Operations ===")

	// Power calculation
	base := 2
	exponent := 5
	fmt.Printf("%d^%d = %d\n", base, exponent, calcPower(base, exponent))

	// Calculate average
	numbers := []int{10, 20, 30, 40, 50}
	avg := calcAverage(numbers)
	fmt.Printf("Average of %v = %.2f\n", numbers, avg)
}

// Basic arithmetic operations
func calcAdd(x, y int) int {
	return x + y
}

func calcSubtract(x, y int) int {
	return x - y
}

func calcMultiply(x, y int) int {
	return x * y
}

func calcDivide(x, y int) int {
	if y == 0 {
		fmt.Println("Error: Cannot divide by zero!")
		return 0
	}
	return x / y
}

func calcModulo(x, y int) int {
	return x % y
}

// Advanced operations
func calcPower(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func calcAverage(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// To run this file:
// go run examples/calculator.go
//
// Expected output:
// === Simple Calculator ===
//
// Numbers: 20 and 8
// Addition:       20 + 8 = 28
// Subtraction:    20 - 8 = 12
// Multiplication: 20 × 8 = 160
// Division:       20 ÷ 8 = 2
// Modulo:         20 % 8 = 4
//
// === Advanced Operations ===
// 2^5 = 32
// Average of [10 20 30 40 50] = 30.00
