package changeLogProc

import (
	"fmt"
	"runtime"
	"sync"
)

// Mylogicalprocs deimonstrates how to create goroutines and
//how the schedulser behames

//This io sthe entry point
func Mylogicalprocs() {
	Myprocessors := runtime.NumCPU()
	runtime.GOMAXPROCS(Myprocessors)

	fmt.Printf("This is the number of Physical Proceesors %v\n: ", Myprocessors)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")
	//Anonymous function creates a goroutine
	go func() {
		//Schedule a call to Done to signal main
		defer wg.Done()
		//Display the alphabet threetimes
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	//Another anonymous function to create second goroutine

	go func() {
		//Schedule a call to Done to signal main
		defer wg.Done()
		//Display the alphabet threetimes
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTerminating now")
}
