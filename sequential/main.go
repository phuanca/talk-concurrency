package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	// SEQ OMIT
	tasks := []string{"task_1", "task_2", "task_3", "task_4", "task_5"}
	for i := range tasks {
		process(tasks[i]) // HL12
	}
	//ENDSEQ OMIT
}

func process(t string) {
	log.Printf("starting %v", t)
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	log.Printf("finish %v", t)
}
