// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

// Add imports.

func main() {

	// Create an unbuffered channel.
	ch := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		defer wg.Done()
		goroutine("A", ch)
	}()

	// Launch the goroutine and handle Done.
	go func() {
		defer wg.Done()
		goroutine("B", ch)
	}()

	// Send a value to start the counting.
	ch <- 0

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(name string, ch chan int) {
	for {

		// Wait to receive a value.
		value, ok := <-ch
		if !ok {

			// If the channel was closed, return.
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", name, value)

		// Terminate when the value is 10.
		if value == 10 {
			close(ch)
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Increment the value and send it
		// over the channel.
		ch <- (value + 1)
	}
}
