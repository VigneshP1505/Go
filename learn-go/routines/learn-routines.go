package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count      int
	volunteers int32
	attendance int64
	mu         sync.Mutex
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker", id)
}

func increment() {
	mu.Lock()
	count++
	mu.Unlock()
}

func incrementWithoutLock(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		count++
	}
}

func incrementWithLock(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func attendanceUpdate() {
	mu.Lock()
	attendance++
	mu.Unlock()
}

func atomicOperators(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&volunteers, 1)
}

// this routine sends data into channel
// main routine receives it
// They synchronize automatically
func channelGoRoutines(ch chan string) {
	ch <- "done"
}

func goRoutineA(ch chan string) {
	ch <- "goRoutine A passed this"
}

func goRoutineB(ch chan string, out chan string) {
	out <- string(<-ch)
}

//scenario 1

func getUser() string {
	time.Sleep(1 * time.Second)
	return "User"
}

func getStripeInfo() string {
	time.Sleep(1 * time.Second)
	return "Info"
}

func getPayments() string {
	time.Sleep(1 * time.Second)
	return "payments"
}

func doWork() {
	start := time.Now()
	user := getUser()
	info := getStripeInfo()
	payments := getPayments()
	fmt.Println(user, info, payments)
	fmt.Println("processing time:", time.Since(start))

	ch := make(chan string)
	routinesStart := time.Now()
	go func() {
		ch <- getUser()
	}()

	go func() {
		ch <- getStripeInfo()
	}()

	go func() {
		ch <- getPayments()
	}()

	for i := 0; i < 3; i++ {
		result := <-ch
		fmt.Println(result)
	}

	fmt.Print(time.Since(routinesStart))
}

// use case 2

func myWorker(id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(1 * time.Second)
	}
}

func processWorker() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go myWorker(1, jobs, &wg)
	go myWorker(2, jobs, &wg)
	go myWorker(3, jobs, &wg)

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

}

// use case 3: backpressure
type Job struct {
	ID int
}

func _inMemoryBackPressureWorker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(2 * time.Second)
	}
}

func _inMemoryTriggerPressure() {
	queueSize := 5
	jobs := make(chan Job, queueSize)

	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go _inMemoryBackPressureWorker(i, jobs, &wg)
	}

	wg.Wait()
}

// goroutines solve a core problem in concurrent programming:
// how multiple goroutines safely communicate and coordinate without corrupting shared state and relying on locks

// with channels, only one goroutine owns data at a time
// no shared memory needed: Goroutine A -> channel -> GoroutineB
func main() {
	_goRoutineLevel8()
}
