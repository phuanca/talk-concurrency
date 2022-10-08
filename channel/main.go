package main

import (
	"log"
	"math/rand"
	"time"
)

type Task string

func main() {
	// Buffered channel.
	ch := make(chan Task, 3)

	// Run n workers.
	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		go runWorker(ch)
	}

	// Send task to workers.
	tasks := []Task{"task_1", "task_2", "task_3", "task_4", "task_5"}

	for _, task := range tasks {
		ch <- task
	}

}

func runWorker(ch chan Task) {
	for {
		task := <-ch
		process(task)
	}
}

func process(t Task) {
	log.Printf("starting %v", t)
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	log.Printf("finish %v", t)
}
