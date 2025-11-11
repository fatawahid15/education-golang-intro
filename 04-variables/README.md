# 04: Variables

This section covers how to declare, initialize, and use variables in Go. It also explains the concepts of zero values and variable scope.

## Code

```go
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
```

### Explanation

- **Declaration with `var`**: The `var` keyword is used to declare one or more variables. You can specify the type, and if you provide an initial value, the type can be inferred.
- **Short Declaration `:=`**: Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with an implicit type. It is not available at the package level.
- **Zero Values**: Variables declared without an explicit initial value are given their *zero value*.
    - `0` for numeric types
    - `false` for the boolean type
    - `""` (the empty string) for strings
    - `nil` for pointers, functions, interfaces, slices, channels, and maps.
- **Scope**: A variable's scope is the region of the code where it is accessible.
    - **Package-level (Global)**: Variables declared outside of any function are visible to all functions in the same package.
    - **Function-level (Local)**: Variables declared inside a function are only visible within that function.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 04-variables
    ```
2.  Run the program:
    ```sh
    go run main.go
    ```

## References

- [Go Tour: Variables](https://go.dev/tour/basics/8)
- [Go Tour: Variables with initializers](https://go.dev/tour/basics/9)
- [Go Tour: Short variable declarations](https://go.dev/tour/basics/10)
- [Go Tour: Zero values](https://go.dev/tour/basics/12)

```