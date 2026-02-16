package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 Welcome to Go Beginner's Toolkit!")
	fmt.Println("=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")

	// 1. Hello World
	helloWorld()

	// 2. Variables and Constants
	variablesDemo()

	// 3. Function with Return Values
	sum := add(5, 3)
	fmt.Printf("\nSum of 5 + 3 = %d\n", sum)

	// 4. Loops and Conditionals
	loopsDemo()

	// 5. Structs
	structDemo()

	// 6. Slices and Arrays
	slicesDemo()

	fmt.Println("\n" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")

	// 7. Example: Hello World
	fmt.Println("\n7️⃣  EXAMPLES FROM EXAMPLES FOLDER")
	RunHelloWorld()

	// 8. Example: Greetings
	fmt.Println()
	RunGreetings()

	// 9. Example: Calculator
	fmt.Println()
	RunCalculator()

	fmt.Println("\n" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")
	fmt.Println("✅ Go Beginner's Toolkit Demo Complete!")
}

// helloWorld prints a simple greeting
func helloWorld() {
	fmt.Println("\n1️⃣  HELLO WORLD")
	fmt.Println("   Hello from Ann's Golang Project 🎉")
	fmt.Printf("   Current time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// variablesDemo demonstrates Go's variable declaration
func variablesDemo() {
	fmt.Println("\n2️⃣  VARIABLES & CONSTANTS")

	// Type inference
	name := "Go Programmer"
	age := 25
	version := 1.21

	fmt.Printf("   Name: %s\n", name)
	fmt.Printf("   Age: %d\n", age)
	fmt.Printf("   Go Version: %.2f\n", version)

	// Constants
	const maxConnections = 100
	fmt.Printf("   Max Connections: %d\n", maxConnections)
}

// add returns the sum of two integers
func add(a, b int) int {
	return a + b
}

// loopsDemo demonstrates loops and conditionals
func loopsDemo() {
	fmt.Println("\n3️⃣  LOOPS & CONDITIONALS")

	// For loop
	fmt.Print("   For loop (1 to 5): ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}

	// If-else
	fmt.Println("\n   Number classification:")
	for num := -1; num <= 2; num++ {
		if num < 0 {
			fmt.Printf("   %d is negative\n", num)
		} else if num == 0 {
			fmt.Printf("   %d is zero\n", num)
		} else {
			fmt.Printf("   %d is positive\n", num)
		}
	}
}

// structDemo demonstrates Go structs
func structDemo() {
	fmt.Println("\n4️⃣  STRUCTS")

	type Person struct {
		Name string
		Age  int
		City string
	}

	// Creating a struct instance
	person := Person{
		Name: "Alice",
		Age:  28,
		City: "Nairobi",
	}

	fmt.Printf("   Person: %s, Age: %d, City: %s\n", person.Name, person.Age, person.City)
}

// slicesDemo demonstrates arrays and slices
func slicesDemo() {
	fmt.Println("\n5️⃣  SLICES & ARRAYS")

	// Array (fixed size)
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Print("   Array: ")
	for _, num := range numbers {
		fmt.Print(num, " ")
	}

	// Slice (dynamic size)
	fruits := []string{"Apple", "Banana", "Orange"}
	fmt.Println("\n   Slice of fruits:")
	for i, fruit := range fruits {
		fmt.Printf("   [%d] %s\n", i, fruit)
	}

	// Append to slice
	fruits = append(fruits, "Mango")
	fmt.Printf("   After append: %v\n", fruits)
}
