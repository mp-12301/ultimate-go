// https://play.golang.org/p/F5DCJTZ6Lm

// Fix the race condition in this program.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// numbers maintains a set of random numbers.
var numbers []int
var m sync.Mutex

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for the application.
func main() {
	// Number of goroutines to use.
	const grs = 3

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create three goroutines to generate random numbers.
	for i := 0; i < grs; i++ {
		go func() {
			random(10)
			wg.Done()
		}()
	}

	// Wait for all the goroutines to finish.
	wg.Wait()

	// Display the set of random numbers.
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}

// random generates random numbers and stores them into a slice.
func random(amount int) {
	// Generate as many random numbers as specified.
	for i := 0; i < amount; i++ {
		n := rand.Intn(100)

		m.Lock()
		{
			numbers = append(numbers, n)
		}
		m.Unlock()
	}
}
