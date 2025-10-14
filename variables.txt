package main

import "fmt"

var middlename string = "Salome" // <= global variable, can be used anywhere

func main() {

	// var age int
	// var name string = "Halo coy"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Salim"
	// type inference ^ allowed variable to be used without declaring any types (:=)

	//Default value
	// Numeric: 0
	// Boolean: false
	// String: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ---- SCOPE

	// fmt.Println(firstName) // <= can't because the firstname isn't on the inside of main func
	fmt.Println(middlename)
	printname()
}

// lifetime of variable only rely on their scope, <= so it won't waste resource too much
// MUST KNOW => function in go only able to be runned in main func scope

func printname() {
	firstName := "Salam"
	fmt.Println(firstName)
	fmt.Println(middlename)
}

// naming variable name have to be meaningful or atleast make it understandable, important for later and make the code maintanable
