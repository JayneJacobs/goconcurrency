package bufferedchannels

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manages a set of resources that can be shared safely by
// multiple goroutines. The resource being managed must implement
// the io.Closer interface.
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

// New creates a poool to manage resources.
// requires: function to allocate a resource, size of pool

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil

}

//Retrieve a resource from the pool
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
		// Provide a new resource if none available
	default:
		log.Println("Acquire: ", "Shared Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	// place a new resource on the pool
	p.m.Lock()
	defer p.m.Unlock()
	//if pool closed , discard resourc
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		log.Println("Release:", "in Queue")

	default:
		log.Println("'Release:'", "Closing")
		r.Close()

	}
}

func (p *Pool) Close() {
	//Secure operation with Release
	p.m.Lock()
	defer p.m.Unlock()

	//If pool is closde do nothing
	if p.closed {
		return
	}
	p.closed = true

	//Close channel before draining of its resources to prevent
	//deadlock
	close(p.resources)
	for r := range p.resources {
		r.Close()

	}
}
