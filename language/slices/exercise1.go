// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var s1 []int

	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
	}

	// Display each value in the slice.
	for i, v := range s1 {
		fmt.Println(i, v)
	}

	// Declare a slice of strings and populate the slice with names.
	s2 := []string{
		"var1",
		"var2",
		"var3",
		"var4",
		"var5",
	}

	// Display each index position and slice value.
	for i, v := range s2 {
		fmt.Println(i, v)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	s3 := s2[1:3]

	fmt.Println("-------------------------")
	// Display each index position and slice values for the new slice.
	for i, v := range s3 {
		fmt.Println(i, v)
	}
}
