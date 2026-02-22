package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Allocate one logical processor for the scheduler to use. default use 1 core in version 1.5 before, use all core at go version 1.5 after
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		lowercase()
		defer wg.Done()
	}()

	go func() {
		uppercase()
		defer wg.Done()
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func lowercase() {
	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

func uppercase() {
	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}
