package main

import "fmt"

func RunHelloWorld() {
	// The simplest Go program - prints to console
	fmt.Println("Hello, World! 👋")

	// You can also use Printf for formatted output
	fmt.Printf("Welcome to Go!\n")

	// Multiple println calls
	name := "Gopher"
	fmt.Printf("Hello, %s!\n", name)
}

// To run this file:
// go run examples/hello-world.go
//
// Expected output:
// Hello, World! 👋
// Welcome to Go!
// Hello, Gopher!
