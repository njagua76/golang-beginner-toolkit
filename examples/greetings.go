package main

import (
	"fmt"
	"strings"
)

func RunGreetings() {
	fmt.Println("=== String Manipulation & Greetings ===")

	// 1. Simple greeting function
	fmt.Println(greet("Alice"))
	fmt.Println(greet("Bob"))
	fmt.Println(greet("Charlie"))

	fmt.Println("\n=== String Operations ===")

	// 2. String operations
	text := "Go Programming"
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(text))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(text))
	fmt.Printf("Length: %d\n", len(text))
	fmt.Printf("Reversed: %s\n", reverseString(text))

	fmt.Println("\n=== String Manipulation ===")

	// 3. String parsing
	words := strings.Split("Go is awesome", " ")
	fmt.Printf("Split 'Go is awesome': %v\n", words)

	// 4. String concatenation
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Printf("Full name: %s\n", fullName)

	// 5. String formatting
	for i := 1; i <= 3; i++ {
		fmt.Printf("Iteration %d: Hello, Go!\n", i)
	}

	fmt.Println("\n=== Personalized Greeting ===")
	displayGreeting("Ann", "software development")
}

// greet returns a personalized greeting message
func greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go! 👋", name)
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// displayGreeting shows detailed greeting information
func displayGreeting(name, interest string) {
	msg := fmt.Sprintf("\n🎉 Welcome, %s!\n", name)
	msg += fmt.Sprintf("   You're interested in: %s\n", interest)
	msg += "   Happy coding!\n"
	fmt.Print(msg)
}

// To run this file:
// go run examples/greetings.go
//
// Expected output:
// === String Manipulation & Greetings ===
//
// Hello, Alice! Welcome to Go! 👋
// Hello, Bob! Welcome to Go! 👋
// Hello, Charlie! Welcome to Go! 👋
//
// === String Operations ===
// Original: Go Programming
// Uppercase: GO PROGRAMMING
// Lowercase: go programming
// Length: 14
// Reversed: gnimmargorP oG
//
// === String Manipulation ===
// Split 'Go is awesome': [Go is awesome]
// Full name: John Doe
// Iteration 1: Hello, Go!
// Iteration 2: Hello, Go!
// Iteration 3: Hello, Go!
//
// === Personalized Greeting ===
//
// 🎉 Welcome, Ann!
//    You're interested in: software development
//    Happy coding!
