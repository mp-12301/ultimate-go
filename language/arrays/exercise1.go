// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var s1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	s2 := [5]string{"vará", "var2", "var3", "var4", "var5"}

	// Assign the populated array to the array of zero values.
	s1 = s2

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for i, v := range s1 {
		fmt.Println(i, v, &s1[i])

		// for j, v := range str {
		// 	fmt.Println(i, v, &str[j])
		// }
	}
}
