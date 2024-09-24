package main

import "fmt"

type Num int

const (
	One Num = iota + 1
	Two
	Three
	Four
)

func miscellaneous() {

	fmt.Println(One)

	greetPtr := greet

	// Call the function via the pointer
	(greetPtr)("Alice") // Dereferencing the pointer

	// Or, you can call it directly since Go allows this shorthand
	greetPtr("Bob")
}

// Define a function
func greet(name string) {
	fmt.Println("Hello,", name)
}
