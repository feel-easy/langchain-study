package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers = 5
	numTasks   = 10
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished task %d\n", id, task)
		wg.Done()
	}
}

func main() {
	tasks := make(chan int, numTasks)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, &wg)
	}

	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		tasks <- i
	}

	close(tasks)

	wg.Wait()
	fmt.Println("All tasks completed.")
}
