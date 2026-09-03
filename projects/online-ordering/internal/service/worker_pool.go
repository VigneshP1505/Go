package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Job struct {
	OrderId uuid.UUID
}

type WorkerPool struct {
	jobs chan Job
}

func (w *WorkerPool) Submit(job Job) {
	w.jobs <- job
}

func (w *WorkerPool) worker(id int) {
	for job := range w.jobs {
		fmt.Print("\n")
		fmt.Printf("Worker %d processing order %s\n", id, job.OrderId)
		time.Sleep(5 * time.Second)
		fmt.Printf("Worker %d done with the processing %s\n", id, job.OrderId)
		fmt.Print()
	}
}

func NewWorkerPool(size int) *WorkerPool {
	wp := &WorkerPool{
		jobs: make(chan Job, 1000),
	}

	for i := range size {
		go wp.worker(i)
	}
	return wp
}
