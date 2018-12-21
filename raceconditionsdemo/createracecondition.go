package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//counter  variable used by all go routines
	counter int
	//wg from the sync to also use for all go routines
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incounter(1)
	go incounter(2)
}

func incounter(id int) {
	//Schedule call to Done to signal main
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//Capture Counter variable
		value := counter
		//Yeild the thread and return to the queue
		runtime.Gosched()
		//increment local value of Coutner variable
		value++
		//Store the value back to counter
		counter = value
	}
	fmt.Printf("I am from %c\n", id)
}

// run with -race flag
// This will sti
