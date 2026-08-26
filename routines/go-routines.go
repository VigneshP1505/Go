package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func _printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(time.Millisecond * 500))
	}
}

func GoRoutine() {
	go _printNumbers()
	time.Sleep(time.Second * 3)
}

func _goRoutineLevel1() {
	fmt.Println("Before:", runtime.NumGoroutine())
	fmt.Println("Before:", runtime.GOARCH)
	fmt.Println("Logical CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}
	fmt.Println("After:", runtime.NumGoroutine())
}

func _goRoutineLevel2() {
	var wg sync.WaitGroup
	fmt.Println("Before:", runtime.NumGoroutine())
	fmt.Println("Before:", runtime.GOARCH)
	fmt.Println("Logical CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println("Done:", id)
		}(i)
	}
	fmt.Println("After:", runtime.NumGoroutine())
	wg.Wait()
}

func cpuIntensiveTask() {
	count := 0

	for n := 2; n < 2_000_000; n++ {
		isPrime := true

		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			count++
		}
	}
}

func _goRoutineLevel3() {
	var wg sync.WaitGroup
	fmt.Println("Before:", runtime.NumGoroutine())
	fmt.Println("Before:", runtime.GOARCH)
	fmt.Println("Logical CPUs:", runtime.NumCPU())
	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cpuIntensiveTask()
			fmt.Println("Done:", id)
		}(i)
	}
	fmt.Println("After:", runtime.NumGoroutine())
	fmt.Println("time taken:", time.Since(start))
}

// preemption allows the 2nd go routine to execute
// from Go 1.14+, asynchronous preemption was introduced
// Suppose scheduler picks up 1st goroutine, older version of Go would never execute the 2nd goroutine as the 1st goroutine never sleeps, never blocks, never, performs I/O, never ends
func _goRoutineLevel4() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello")
	}()

	wg.Wait()
}

// 6 go routines - 5 workers, 1 main routine
// 2 processors
// At any instance simultaneously only two workers can be run
// other workers are in Runnable state, not Running
// time.Sleep() blocks the go orutines, not the processor
func work(id int) {
	fmt.Println("Worker", id, "started")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker", id, "finished")
}
func _goRoutineLevel5() {
	runtime.GOMAXPROCS(2)
	for i := 1; i <= 5; i++ {
		go work(i)
	}
	time.Sleep(3 * time.Second)
}

// the go scheduler promises to be efficient but does not guarantee FIFO order

// When G1 executes <-ch, it enters blocked or Waiting state
// P1 is never blocked, only goroutines are blocked
// M1 is never blocked
// G2 is run after G1 is blocked as G2 is in Runnable state
// G1 running, sends value 42, G1 becomes Waiting -> Runnable. G2 can run longer and G1 runs only when it is done or is blocked or is Waiting
func _goRoutineLevel6() {
	runtime.GOMAXPROCS(1)
	ch := make(chan int)

	go func() {
		fmt.Println("G1:Waiting for value...")
		x := <-ch
		fmt.Println("G1 received:", x)
	}()

	go func() {
		fmt.Println("G2: Sleeping...")
		time.Sleep(2 * time.Second)
		fmt.Println("G2: Sending value")
		ch <- 42
	}()

	time.Sleep(3 * time.Second)
}

// work stealing
// GOMAXPROCS = 2

// P1 Run Queue:
// G1
// G2
// G3
// G4
// G5
// G6

// P2 Run Queue:
// (empty)

// M2 is a OS thread, it can or cannot be idle. It might be running some other instruction
// Since P2 local queue is empty, it steals some goroutines from P1
// P2 cannot steal currently running routine. It can only steal Runnable goroutine
// P1->M1->G1, P2->M2->any of the other runnable goroutines

// P Handoff: If a goroutine makes a blocking system call, the runtime detaches the processor (P) from that blocked thread (M) and attaches it to a fresh or waiting thread so other tasks keep moving.
// OS-level blocking - blocking syscall certain file I/O, device I/O
// Go achieves parallelism automatically on multi-core machines. Its internal scheduler (the M:N scheduler) distributes active goroutines across multiple operating system threads and CPU cores.
