package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {

	const grs = 2

	wg := sync.WaitGroup{}
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2000; count++ { // setting count < 2000 The competition is quite fierce
				mu.Lock() // plus mutex to ensure the atomicity of execution
				value := counter

				// Yield the thread and be placed back in queue
				//runtime.Gosched()

				value++
				counter = value
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
}
