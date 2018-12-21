package main

import (
	"goConcurrency/unbufferedchannels"
	"log"
	"sync"
	"time"
)

//Demo work pacakage to guarantee work is done

var names = []string{
	"Jayne",
	"Ali",
	"Steve",
	"Pranay",
	"Dave",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	//Create a work pool with 2 goroutines
	p := unbufferedchannels.New(2)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			//Create namePrinter and provide the name
			np := namePrinter{
				name: name,
			}

			go func() {
				//Submit task to be worked on When Runtask
				//returns
				p.Run(&np)
				wg.Done()
			}()

		}

	}
	wg.Wait()

	//Shutdown work pool after work complets
	p.Shutdown()
}
