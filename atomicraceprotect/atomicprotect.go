package atomicraceprotect

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	//counter fariable incremented by all goroutines
	counter int64

	//wg to siganal program to finish
	wg1 sync.WaitGroup
)

func AtomicProtect() {
	//Add 2 Wait Groups
	wg1.Add(2)
	//Create goroutines
	go incCounter(1)
	go incCounter(2)
	//Wait for goroutines to finish
	wg1.Wait()
	//Display value
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg1.Done()

	for count := 0; count < 2; count++ {
		//Safely Add one to counter
		atomic.AddInt64(&counter, 1)
		//Yield the thread and go back to the queue
		runtime.Gosched()

	}
}
