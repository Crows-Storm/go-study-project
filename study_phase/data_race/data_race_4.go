package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	data []string

	rwMutex   sync.RWMutex
	readCount int64
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	for i := 0; i < 8; i++ {
		go func(i int) {
			for {
				reader(i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Program Complete: ", data)
}

// rwMutex.Lock() Wait for RLock to be released before locking and writing data
func writer(i int) {
	rwMutex.Lock()
	{
		rc := atomic.LoadInt64(&readCount)

		fmt.Printf("****> : Performing Write : RCount[%d]\n", rc) // rc always 0, because rwMutex.Lock() need wait RLock released
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

func reader(id int) {
	rwMutex.RLock()
	{
		rc := atomic.AddInt64(&readCount, 1)

		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc)

		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
}
