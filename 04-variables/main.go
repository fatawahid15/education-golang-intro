package main

import "fmt"

// middleName is a package-level (global) variable.
// It can be accessed by any function in the 'main' package.
var middleName string = "Salome"

func main() {
	// Using 'var' to declare and initialize a variable.
	// The type is explicitly defined.
	var age int = 30
	fmt.Println("Age:", age)

	// Type inference with 'var'.
	// Go infers that the type of 'name' is string.
	var name = "Jane"
	fmt.Println("Name:", name)

	// Short variable declaration using ':='.
	// This is the most common way to declare and initialize variables inside functions.
	// The type is inferred.
	count := 10
	fmt.Println("Count:", count)

	// --- Zero Values ---
	// When a variable is declared but not initialized, it is given a "zero value".
	var (
		defaultInt    int
		defaultFloat  float64
		defaultBool   bool
		defaultString string
	)
	fmt.Println("--- Zero Values ---")
	fmt.Printf("Default Integer: %d\n", defaultInt)
	fmt.Printf("Default Float: %f\n", defaultFloat)
	fmt.Printf("Default Boolean: %t\n", defaultBool)
	fmt.Printf("Default String: '%s'\n", defaultString)
	fmt.Println("-------------------")

	// --- Scope ---
	// 'middleName' is accessible here because it's declared at the package level.
	fmt.Println("Middle Name (from main):", middleName)

	// Call another function to demonstrate scope.
	printName()
}

func printName() {
	// 'firstName' is local to the printName function.
	// It cannot be accessed from the 'main' function.
	firstName := "Salam"
	fmt.Println("First Name (from printName):", firstName)

	// 'middleName' is also accessible here.
	fmt.Println("Middle Name (from printName):", middleName)
}
