package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	track := make(chan int)

	wg.Add(1)

	go Runner(track)

	track <- 1

	wg.Wait()
}

func Runner(track chan int) {
	const maxExchanges = 4

	var exchange int

	baton := <-track

	fmt.Printf("Runner %d Running With Baton\n", baton)

	if baton < maxExchanges {
		exchange = baton + 1
		fmt.Printf("Runner %d To The Line\n", exchange)
		go Runner(track)
	}

	time.Sleep(100 * time.Millisecond)

	if baton == maxExchanges {
		fmt.Printf("Runner %d Finished, Race Over\n", baton)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", baton, exchange)

	track <- exchange
}
