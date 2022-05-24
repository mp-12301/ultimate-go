// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var s string
	var i int
	var f bool

	// Display the value of those variables.
	fmt.Printf("Value of s \t %v\n", s)
	fmt.Printf("Value of i \t %v\n", i)
	fmt.Printf("Value of f \t %v\n", f)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	s1 := "3.14"
	i1 := 10
	f1 := true

	// Display the value of those variables.
	fmt.Printf("Value of s1 \t %v\n", s1)
	fmt.Printf("Value of i1 \t %v\n", i1)
	fmt.Printf("Value of f1 \t %v\n", f1)

	// Perform a type conversion.
	var pi float32
	pi = float32(3.14)

	// Display the value of that variable.
	fmt.Printf("value of pi \t %T %v\n", pi, pi)
	fmt.Printf("value of literal pi \t %T %v\n", 3.14, 3.14)
}
