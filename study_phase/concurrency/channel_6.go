package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

const timeoutSeconds = 3 * time.Second

var (
	sigChan = make(chan os.Signal, 1)

	timeout = time.After(timeoutSeconds)

	complete = make(chan error)

	shutdown = make(chan struct{})
)

func main() {
	log.Println("Starting Process")

	signal.Notify(sigChan, os.Interrupt)

	// Launch the process.
	log.Println("Launching Processors")

	go processor(complete)

ControlLoop:
	for {
		select {
		case <-sigChan:
			log.Println("OS INTERRUPT")
			close(shutdown)
			sigChan = nil
		case <-timeout:
			log.Println("Timeout - Killing Program")
			os.Exit(1)
		case err := <-complete:
			log.Printf("Task Completed: Error[%s]", err)
			break ControlLoop
		}
	}

	log.Println("Process Ended")
}

func processor(complete chan<- error) {
	log.Println("Starting processor")
	var err error
	defer func() {
		if r := recover(); r != nil {
			log.Println("Processor - Panic", r)
		}
		complete <- err
	}()

	err = doWork()

	log.Println("Processor - Completed")
}

func doWork() error {

	log.Println("Processor - Task 1")
	time.Sleep(2 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 3")
	time.Sleep(1 * time.Second)

	return nil
}

func checkShutdown() bool {
	select {
	case <-shutdown:
		log.Println("checkShutdown - Shutdown Early")
		return true
	default:
		return false
	}
}
