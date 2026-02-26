package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	UserID string
}

func main() {

	deadLine := time.Now().Add(150 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), deadLine)
	defer cancel()
	ch := make(chan data, 1)

	go func() {
		time.Sleep(200 * time.Millisecond) // time out so ctx done
		//time.Sleep(100 * time.Millisecond)	// not time out, so ch msg is consumer

		ch <- data{"123"}
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
