package main

import (
	"fmt"
	"goConcurrency/mutexraceprotect"
	"sync"
)

//used to wait for the program to end
var wg sync.WaitGroup

//main is the entry point

func main() {
	//1 Logical processor for the scheduler
	// runtime.GOMAXPROCS(1)
	// //add two to go routines to the wait group
	// wg.Add(2)
	// // Create Goroutines
	// fmt.Println("Create Gorouteines")

	// go printPrime("A")
	// go printPrime("B")

	// fmt.Println("Waiting to Finish")
	// wg.Wait()

	// fmt.Println("Terminating the Create Go Routines process")
	// createfile.MyFileMake("./", "waitgroupout.txt")
	// changeLogProc.Mylogicalprocs()
	// fmt.Println("Start with 2 processors")
	// changeLogProc.Mytwoprocessors()
	// fmt.Println("Start Atomic Protect Demo")
	// atomicraceprotect.AtomicProtect()
	// atomicraceprotect.Mystoreandload()
	fmt.Println("Start Mutex Protect Demo")
	mutexraceprotect.Mutexsync()
	// fmt.Println("Unbuffered Channels")
	// unbufferedchannels.MyUnbuffered()
	// bufferedchannels.Mybufferedchannel()
}

//printPrime displays prime numbers fo the first 5000 numbers
/*Because of the random nature of the program and the Go scheduler,
the output for this program will be different every time you run it.
 But the use of all four goroutines to process work from the
 buffered channel won’t change. You can see from the output how each
  goroutine is receiving work distributed from the channel.*/

// func printPrime(prefix string) {
// 	//Schedule a cal to Done to signal main
// 	defer wg.Done()
// next:
// 	for outer := 2; outer < 5000; outer++ {
// 		for inner := 2; inner < outer; inner++ {
// 			if outer%inner == 0 {
// 				continue next
// 			}
// 		}
// 		fmt.Printf("%s:%d\n", prefix, outer)
// 	}
// 	fmt.Println("Completed", prefix)
// }
