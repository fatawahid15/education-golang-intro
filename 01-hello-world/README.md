# 01: Hello, World!

This section covers the classic "Hello, World!" program in Go. It's the first step in learning any new programming language.

## Code

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

### Explanation

-   **`package main`**: In Go, every program is part of a package. The `main` package is special; it tells the Go compiler that the package should compile as an executable program instead of a shared library. The `main` function in the `main` package is the entry point of the program.

-   **`import "fmt"`**: This line imports the `fmt` package, which is one of Go's standard library packages. The `fmt` package implements formatted I/O (input/output) with functions analogous to C's `printf` and `scanf`.

-   **`func main()`**: This is the main function where the program execution begins.

-   **`fmt.Println("Hello, World!")`**: This is a statement that calls the `Println` function from the `fmt` package. `Println` prints a line of text to the standard output (usually your terminal) and adds a newline character at the end.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 01-hello-world
    ```
2.  Run the program using the `go run` command:
    ```sh
    go run main.go
    ```
3.  You should see the following output:
    ```
    Hello, World!
    ```

## References

-   [Go Tour: Hello, World](https://go.dev/tour/hello/1)
-   [Go Packages](https://go.dev/tour/basics/1)
-   [The `fmt` package documentation](https://pkg.go.dev/fmt)
