package concurrencyPatterns

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	//interupt channel reports signal from os
	interrupt chan os.Signal
	// complete channel report processing is done
	complete chan error

	//timeout reports time has run out
	timeout <-chan time.Time
	//task hods functions
	// executed synchronously in order
	tasks []func(int)
}

// ErrTimeout returned when value received on timeout
var ErrTimeout = errors.New("received Timeout")

//ErrINterrrupt returned when event from OS receied.
var ErrInterrupt = errors.New("received interrupt")

//New returns a new ready to use Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add attaches tasks to Runner. functions that take an int ID.

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//Start runs tasks and monitors channel events
// takes and int ID
func (r *Runner) Start() error {
	//receive all interupt signals
	signal.Notify(r.interrupt, os.Interrupt)
	// Run the tasks on a go routene
	go func() {
		r.complete <- r.run()
	}()
	select {
	//Signaled processing is done
	case err := <-r.complete:
		return err
		//Signal timeout
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		//Check for interrupt
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		//execute task
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	//signal ineterupt
	case <-r.interrupt:
		//Stop receiving signals
		signal.Stop(r.interrupt)
		return true

		//Continue running default
	default:
		return false
	}
}
