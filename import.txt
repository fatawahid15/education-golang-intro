package main

// import (
// 	"fmt"
// 	"net/http"
// )

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Go Standard Library!")

	resp, err := foo.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
