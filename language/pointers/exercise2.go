// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

// Add imports.

// Declare a type named user.
type user struct {
	name string
	age  int
}

// Create a function that changes the value of one of the user fields.
func funcName(u *user) {
	u.age++
}

func main() {

	// Create a variable of type user and initialize each field.
	u := user{
		name: "John doe",
		age:  20,
	}

	// Display the value of the variable.
	fmt.Printf("%v \n", u)

	// Share the variable with the function you declared above.
	funcName(&u)

	// Display the value of the variable.
	fmt.Printf("%v \n", u)
}
