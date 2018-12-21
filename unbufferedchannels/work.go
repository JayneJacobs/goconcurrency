package unbufferedchannels

import (
	"sync"
)

/*work package is to show how you can use an unbuffered channel
to create a pool of goroutines that will perform and control
the amount of work that gets done concurrently. This is a better
 approach than using a buffered channel of some arbitrary
 static size that acts as a queue of work and throwing
 a bunch of goroutines at it.

 Unbuffered channels provide a guarantee that data has been
 exchanged between two goroutines. This approach of using an unbuffered
 channel allows the user to know when the pool is performing the work,
 and the channel pushes back when it can’t accept any more work because
 it’s busy. No work is ever lost or stuck in a queue that has no guarantee
 it will ever be worked on.*/

type Worker interface {
	Task()
}

//Pool provides a pool of goroutines that execute any Work tasks submitted
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

//New createes a new work pool
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

//Run submits work to the pools
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//Shutdown wiats for goroutines to shutdown
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
