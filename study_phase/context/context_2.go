package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		time.Sleep(50 * time.Millisecond)
		//time.Sleep(150 * time.Millisecond)	// time out, in to case <-time.After(100 * time.Millisecond)
		cancel()
	}()

	select {
	case <-time.After(100 * time.Millisecond): // return <-chan Time the channel
		fmt.Println("moving on")
	case <-ctx.Done(): // return <-chan struct{} the channel
		fmt.Println("work complete")
	}

}
