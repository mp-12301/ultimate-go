// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

import "fmt"

// Add imports.

const (
	// Declare a constant named server of kind string and assign a value.
	server = "foobar"
	// Declare a constant named port of type integer and assign a value.
	port int = 9999
)

func main() {

	// Display the value of both server and port.
	fmt.Printf("server %v\n", server)
	fmt.Printf("port %v\n", port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	result := 2 * 2.22

	// Display the value of the variable.
	fmt.Printf("result %v\n", result)
}
