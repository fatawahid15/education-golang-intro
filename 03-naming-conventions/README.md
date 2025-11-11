# 03: Naming Conventions

This section explains the idiomatic naming conventions used in Go. Effective naming is crucial for writing readable and maintainable code.

## Code

```go
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
```

### Explanation

Go's naming conventions are simple and serve to define the visibility of identifiers (whether they are public or private).

-   **Exported vs. Unexported (Public vs. Private)**
    -   If a name (like a variable, type, or function) starts with a capital letter, it is **exported** (public). Exported identifiers can be accessed from other packages.
    -   If a name starts with a lowercase letter, it is **unexported** (private). Unexported identifiers are only accessible within their own package.

-   **`PascalCase` or `camelCase`? (MixedCaps)**
    -   Go prefers using `MixedCaps` or `mixedCaps` rather than using underscores (`snake_case`).
    -   **`PascalCase`** (e.g., `UserInfo`, `CalculateTotal`) is used for exported identifiers.
    -   **`camelCase`** (e.g., `userID`, `printUserDetails`) is used for unexported identifiers.

-   **Constants**
    -   Constants should be named using `MixedCaps` or `mixedCaps` just like variables. For example, `const MaxRetries = 5` or `const maxRetries = 5`.
    -   Using all uppercase letters (e.g., `MAX_RETRIES`) is not a common Go idiom, but it is not forbidden, especially for acronyms.

-   **Acronyms**
    -   Acronyms like `URL`, `ID`, `HTTP`, `API` should be written in all uppercase letters, both in exported and unexported names. For example: `apiClient`, `ServeHTTP`, `UserID`.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 03-naming-conventions
    ```
2.  Run the program:
    ```sh
    go run main.go
    ```
3.  You should see the following output:
    ```
    User ID: 1
    First Name: John
    Max Retries: 5
    ```

## References

-   [Go Tour: Exported names](https://go.dev/tour/basics/3)
-   [Effective Go: Names](https://go.dev/doc/effective_go#names)
-   [Go Code Review Comments: Mixed Caps](https://github.com/golang/go/wiki/CodeReviewComments#mixed-caps)

```