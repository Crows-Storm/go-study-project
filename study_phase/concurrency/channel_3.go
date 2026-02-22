package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		player("Crow", court)
		wg.Done()
	}()

	go func() {
		player("Andrew", court)
		wg.Done()
	}()

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d Striking %d \n ", name, ball, n)
		ball++

		court <- ball
	}
}
