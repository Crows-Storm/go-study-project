package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {

	const grs = 2

	wg := sync.WaitGroup{}
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2000; count++ { // setting count < 2000 The competition is quite fierce
				atomic.AddInt64(&counter, 1) // Atomic operations solve the data race problem caused by multiple reads and writes

				runtime.Gosched()

			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
}
