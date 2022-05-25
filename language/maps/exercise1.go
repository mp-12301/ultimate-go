// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

// Add imports.

func main() {

	// Declare and make a map of integer type values.
	n1 := make(map[string]int)

	// Initialize some data into the map.
	n1["var1"] = 1
	n1["var2"] = 2
	n1["var3"] = 3
	n1["var4"] = 4
	n1["var5"] = 5

	// Display each key/value pair.
	for k, v := range n1 {
		fmt.Println(k, v)
	}
}
