package atomicraceprotect

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// This demonstrates the use of teh atomic package functions Store and Load for safe acces
// to numeric types

var (
	//shutdown is a flac to alert goroutines to shutdown
	shutdown int64
	//wg waits for a program to fininsh
	//this is declared in the atomicprotect
	wg sync.WaitGroup
)

func Mystoreandload() {
	//Add count of for each gorouting
	wg.Add(2)

	//create goroutines
	go doWork("A")
	go doWork("B")

	//Give goroutines time to run
	time.Sleep(1 * time.Second)
	//Safely flag time to shutdown
	fmt.Println("Shutdown nwo")
	atomic.StoreInt64(&shutdown, 1)
	//Wait for goroutines to finish
	wg.Wait()

}

func doWork(name string) {
	defer wg.Done()
	//endlessfor loop
	for {
		fmt.Printf("Doing %s Work \n", name)
		time.Sleep(250 * time.Millisecond)

		//Do we need to shutdown
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
