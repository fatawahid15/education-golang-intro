package main

import "fmt"

// User is an example of a struct with exported fields.
// Exported types and fields start with a capital letter (PascalCase).
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// a private constant (not exported)
const defaultPermissions = "user"

// MAX_RETRIES is an example of a constant. In Go, constants are often
// written in PascalCase or camelCase like other variables, but using
// all caps is acceptable for acronyms or to signal a very stable constant.
const MAX_RETRIES = 5

func main() {
	// Variables are typically written in camelCase.
	var userID int = 1
	firstName := "John"

	fmt.Println("User ID:", userID)
	fmt.Println("First Name:", firstName)
	fmt.Println("Max Retries:", MAX_RETRIES)
}

// printUserDetails is an example of a private function (not exported)
// because it starts with a lowercase letter.
func printUserDetails(user User) {
	fmt.Printf("User: %s %s (ID: %d)\n", user.FirstName, user.LastName, user.ID)
}

// CalculateTotal is an example of an exported function.
func CalculateTotal() {
	// function logic here
}

