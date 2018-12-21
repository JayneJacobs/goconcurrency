package changeLogProc

import (
	"fmt"
	"runtime"
	"sync"
)

// Mytwoprocessors demonstrates how to create
//goroutines and how the scheduler behaves with 2 logical processors
func Mytwoprocessors() {
	//Allocate two logical processors for the runtime to use.
	runtime.GOMAXPROCS(2)
	//wg used to wait for prgram to finish
	//Add count of 2, one for each goroutine

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")
	//Declare an anonymous function and create a go routine
	go func() {
		//Schedue the call to Done and defer till main the function is done
		defer wg.Done()
		//Display alphabet 3 times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	go func() {
		//Schedue the call to Done and defer till main the function is done
		defer wg.Done()
		//Display alphabet 3 times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)

			}
		}
	}()
	// Wait for goroutines to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTerminating Process")
}
