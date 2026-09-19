package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow Greet", phrase)
	doneChan <- true
}

func Routines() {
	// go greet("Nice to meet you?")
	// go greet("How are you!")
	done := make(chan bool)
	go slowGreet("How...are...you!", done)
	go greet("I hope you are okay!", done)
	<-done
	<-done

	channels := make([]chan bool, 4)
	for i := range channels {
		channels[i] = make(chan bool)
	}
	go greet("This is message 1", channels[0])
	go greet("This is message 2", channels[1])
	go greet("This is message 3", channels[2])
	go greet("This is message 4", channels[3])

	for _, channel := range channels {
		fmt.Println(<-channel)
	}
}

//go can handle 100k+ concurrent goroutines without much memory
//go routines are lightweight managed threads. They run in parallel on multiple OS threads. They are multiplexed on OS threads by go scheduler
//Channels allow goroutines to communicate without shared memory
