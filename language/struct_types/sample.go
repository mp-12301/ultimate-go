// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// Add imports.

// Add user type and provide comment.

type user struct {
	name  string
	email string
	age   int
}

func main() {
	// Declare variable of type user and init using a struct literal.
	u := user{
		name:  "John",
		email: "john@doe.com",
		age:   35,
	}

	// Display the field values.
	fmt.Printf("value of user %v\n", u)

	// Declare a variable using an anonymous struct.
	anon := struct {
		name  string
		email string
		age   int
	}{
		name:  "Foo",
		email: "Foo@Bar.com",
		age:   26,
	}

	// Display the field values.
	fmt.Printf("value of user anon %v %T", anon, anon)
}
