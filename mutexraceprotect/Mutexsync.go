package mutexraceprotect

import (
	"fmt"
	"runtime"
	"sync"
)

//Demo how to use a mutex to define critical sections of code for sync

var (
	//counter variable shared by goroutines
	counter int
	//wait gourp used to signal complete
	wg sync.WaitGroup

	//mutex todefine section of code to protect
	mutex sync.Mutex
)

func main() {
	//Add wg for each goroutine
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	//Wait for goroutines to finish
	wg.Wait()
	fmt.Printf("Final Counter: %d\\n", counter)
}

func incCounter(id int) {
	//Allows one goroutine at a time
	defer wg.Done()
	mutex.Lock()
	{
		value := counter
		//Yield the thread and place back in the queue
		runtime.Gosched()
		//Increment local value of counter
		value++

		//Store value back to counter
		counter = value
	}
	mutex.Unlock()
}
