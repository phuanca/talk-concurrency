package main

import (
	"log"
	"time"
)

// CHAN OMIT
type Task string

func main() {
	// Buffered channel.
	ch := make(chan Task, 4) // HL
	// Run n=4 workers.
	for i := 0; i < 4; i++ {
		go runWorker(ch) // HL
	}
	// Send task to workers.
	tasks := []Task{"task_1", "task_2", "task_3", "task_4", "task_5"}
	for _, task := range tasks {
		ch <- task // HL
	}
	log.Printf("main is waiting...")
	time.Sleep(6 * time.Second)
}

func runWorker(ch chan Task) { // HL
	for {
		task := <-ch  // HL
		process(task) // HL
	}
}

// ENDCHAN OMIT

func process(t Task) {
	log.Printf("starting %v", t)
	// randTime := time.Duration(rand.Intn(5000)) * time.Millisecond
	randTime := 5 * time.Second
	time.Sleep(randTime)
	log.Printf("finish %v in: %v", t, randTime)
}
