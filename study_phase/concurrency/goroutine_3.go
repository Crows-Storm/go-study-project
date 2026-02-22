package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// when the number of cores is > 1, single-core competition will turn into multicore concurrency
	runtime.GOMAXPROCS(2) // created 2 core P
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() { // first m
		for count := 0; count < 3; count++ { // first goroutine
			for r := 'a'; r <= 'z'; r++ {
				fmt.Printf("%c ", r)
			}
		}

		wg.Done()
	}()

	go func() { // second m
		for count := 0; count < 3; count++ { // second goroutine
			for r := 'A'; r <= 'Z'; r++ {
				fmt.Printf("%c ", r)
			}
		}

		wg.Done() // wg - 1
	}()

	fmt.Println("Waiting To Finish")
	defer wg.Wait() // wait wg = 0

	fmt.Println("\nTerminating Program")
}
