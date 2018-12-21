package unbufferedchannels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Examples that use an unbuffered channel to synchronize
//the exchange of data between two goroutines.

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func MyUnbuffered() {
	court := make(chan int)

	//Add a wg for each goroutine
	wg.Add(2)
	//Launch go routines
	go player("Jayne", court)
	go player("Ali", court)
	//Start the set
	court <- 1
	//Wait for game to finis
	wg.Wait()
}

func player(name string, court chan int) {
	//Schedule the call to Done
	defer wg.Done()
	for {
		//Wait for the ball to be hit
		ball, ok := <-court
		if !ok {
			//if channel closes we win
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			//Close channel to signal we lost
			close(court)
			return
		}
		//Display and increment the hit count
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		//hit the ball back
		court <- ball
	}
}
