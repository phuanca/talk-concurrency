package main

import (
	"log"
	"math/rand"
	"time"
)

func main() { // HL12

	// CONCUR OMIT
	tasks := []string{"task_1", "task_2", "task_3", "task_4", "task_5"}
	for i := range tasks {
		go process(tasks[i]) // HL
	}
	// Need to wait?
	time.Sleep(6 * time.Second)
	//ENDCONCUR OMIT
}

// PROC OMIT
func process(t string) { // HL
	log.Printf("starting %v", t)
	randTime := time.Duration(rand.Intn(5000)) * time.Millisecond
	time.Sleep(randTime)
	log.Printf("finish %v in: %v", t, randTime)
}

//ENDPROC OMIT
