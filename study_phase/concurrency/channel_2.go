package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("\n=> Double signal\n")
	signalAck()

	fmt.Printf("\n=> Select and receive\n")
	selectRecv()

	fmt.Printf("\n=> Select and send\n")
	selectSend()

	fmt.Printf("\n=> Select and drop\n")
	selectDrop()
}

func signalAck() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
		ch <- "ok done"
	}()

	ch <- "do this"
	fmt.Println(<-ch)
}

func closeRange() {
	// This is a buffered channel of 5.
	ch := make(chan int, 5)

	// Populate with value
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// Close the channel.
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func selectRecv() {
	ch := make(chan string)

	// Wait for some amount of time and perform a send.
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "work"
	}()

	select {
	// Used to handle multiple channels
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

}

func selectSend() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Println(<-ch)
	}()

	select {
	case ch <- "work":
		fmt.Println("send work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

}

func selectDrop() {
	ch := make(chan int, 5)
	go func() {
		// We are in the receive loop waiting for data to work on.
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
			fmt.Println("send work", i)
		default:
			fmt.Println("drop", i)
		}
	}

	close(ch)
}
