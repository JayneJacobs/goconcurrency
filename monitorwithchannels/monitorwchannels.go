package main

import (
	"goConcurrency/concurrencyPatterns"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work. ")

	//Create timer value for run
	r := concurrencyPatterns.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case concurrencyPatterns.ErrTimeout:
			log.Println("Terminating due to timeout. ")
			os.Exit(1)
		case concurrencyPatterns.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}

	}
	log.Println("Process ended")
}

//create Tasks and sleep to simulate simulate timeout
//push control-c to test interrupt
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
