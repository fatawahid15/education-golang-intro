# 02: Imports

This section explains how to import packages in Go, including how to alias them.

## Code

```go
package main

// A standard import block
// import (
// 	"fmt"
// 	"net/http"
// )

// An import block with an alias
import (
	"fmt"
	foo "net/http" // 'foo' is an alias for the "net/http" package
)

func main() {
	fmt.Println("Go Standard Library!")

	// We use 'foo' to access functions from the "net/http" package
	resp, err := foo.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // Closes the response body when the function exits

	fmt.Println("Response Status:", resp.Status)
}
```

### Explanation

-   **`import (...)`**: The `import` declaration is used to import packages. You can import multiple packages by listing them inside parentheses.
-   **Aliasing Imports**: You can provide an alias for an imported package. In the example, `foo` is an alias for the `net/http` package. This is useful when you have packages with the same name or when a package name is too long.
-   **`net/http` package**: This is a standard library package that provides HTTP client and server implementations. In this example, we use `http.Get` to send an HTTP GET request to a URL.
-   **`defer resp.Body.Close()`**: The `defer` statement schedules a function call (in this case, `resp.Body.Close()`) to be run immediately before the function (`main`) returns. It's a common idiom in Go to ensure resources like network connections or files are closed.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 02-imports
    ```
2.  Run the program:
    ```sh
    go run main.go
    ```
3.  You should see an output similar to this:
    ```
    Go Standard Library!
    Response Status: 200 OK
    ```

## References

-   [Go Tour: Imports](https://go.dev/tour/basics/2)
-   [Go Tour: Packages](https://go.dev/tour/basics/1)
-   [The `net/http` package documentation](https://pkg.go.dev/net/http)
-   [A Tour of Go: Defer](https://go.dev/tour/flowcontrol/12)
