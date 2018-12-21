package bufferedchannels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10 //Amount of work to process

)

var wg sync.WaitGroup

func init() {
	//Seed random number generator
	rand.Seed(time.Now().Unix())
}

func Mybufferedchannel() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//Add work to do
	for post := 0; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}
	//Close the channel  goroutines will quit
	// when finished
	close(tasks)
	//Wait to finish work
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	//Report a return
	defer wg.Done()

	for {
		//Wait for work to be assigned
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		//Display start of work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		//display Finish
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}

}
