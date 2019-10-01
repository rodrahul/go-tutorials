package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)
	var mutex sync.Mutex

	var readOps, writeOps uint64

	// Create 100 go routines to read from the map
	for i := 0; i < 100; i++ {
		go func() {
			for {
				total := 0
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Create 10 go routines to write to the map
	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)

				time.Sleep(time.Millisecond)
			}

		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("readOps:", readOpsFinal)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
