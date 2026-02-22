package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\n=> Basics of a send and receive\n")
	basicSendRecv()

	fmt.Printf("\n=> Close a channel to signal an event\n")
	signalClose()
}

func basicSendRecv() {
	// This is an unbuffered channel.
	ch := make(chan string)

	go func() {
		ch <- "hello"
	}()

	value, ok := <-ch
	fmt.Println(value, ok)
}

func signalClose() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("signal event")
		close(ch)
	}()

	value, ok := <-ch
	fmt.Println(value, ok)

	fmt.Println("event received")
}
