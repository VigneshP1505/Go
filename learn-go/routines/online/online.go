package online

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID int
}

var (
	totalProcessed int
	mu             sync.Mutex
)

func worker(id int, jobs <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range jobs {
		fmt.Printf("Worker %d processing order %d\n", id, order.ID)
		time.Sleep(200 * time.Millisecond) // simulate work

		mu.Lock()
		totalProcessed++
		mu.Unlock()
	}
}

func OrderHandler() {
	jobs := make(chan Order, 5)
	var wg sync.WaitGroup

	// Start 3 workers (goroutines)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send 10 jobs
	for i := 1; i <= 10; i++ {
		jobs <- Order{ID: i}
	}
	close(jobs) // no more jobs

	wg.Wait() // wait for all workers to finish

	fmt.Println("Total processed:", totalProcessed)
}

// Waitgroups are counters waiting for goroutines to finish
// They block the execution until all goroutines are done - counter reaches 0
func workers(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker with id %d is done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workers(i, &wg)
	}
	wg.Wait()
	fmt.Printf("done with all workers")
}

// Channels are typed pipes used for communication between goroutines safely and synchronously
func Channels() {
	channel := make(chan string)
	go func() {
		channel <- "data 1"
		channel <- "data 2"
		channel <- "data 3"
		close(channel)
	}()
	for msg := range channel {
		fmt.Println(msg)
	}

	channels := make(chan string, 3)
	go func() {
		channels <- "data 1"
		channels <- "data 2"
		channels <- "data 3"
		close(channels)
	}()
	for msg := range channels {
		fmt.Println(msg)
	}

	unBufferedChannel := make(chan string)
	go func() {
		unBufferedChannel <- "data 1"
		fmt.Println("This is blocked")
	}()
	fmt.Println("Trying to start")
	time.Sleep(1 * time.Second)
	var val string = <-unBufferedChannel
	fmt.Printf("val:%s\n", val)
	fmt.Println("Done")

	fmt.Println("---------------")

	bufferedChannel := make(chan string, 3)
	bufferedChannel <- "data 1"
	bufferedChannel <- "data 2"
	go func() {
		bufferedChannel <- "data 3"
		fmt.Println("This will not be blocked")
		bufferedChannel <- "data 4"
		fmt.Println("This will be blocked")
	}()
	for msg := range bufferedChannel {
		fmt.Println(msg)
	}
}

// Mutex- mutex is a lock to prevent shared memory race conditions
var count int = 0
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	count++
	mutex.Unlock()
}

func Mutex() {
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Printf("count:%d\n", count)
}

//select lets a goroutine wait on multiple channel operations at the same time
