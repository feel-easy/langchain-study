package main

import (
	"sync"
	"time"
)

type Worker struct {
	ID int
	Ch chan int
	WG *sync.WaitGroup
}

func NewWorker(id int, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID: id,
		Ch: make(chan int, 10),
		WG: wg,
	}
}

func (w *Worker) Start() {
	for i := range w.Ch {
		time.Sleep(time.Second)
		println("id", w.ID, "ret", i*2)
	}
	w.WG.Done()
}

func (w *Worker) Stop() {
	close(w.Ch)
}

func main() {
	workers := make([]*Worker, 0, 5)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		wk := NewWorker(i, wg)
		workers = append(workers, wk)
		go wk.Start()
	}
	for i := 0; i < 10; i++ {
		workers[i%5].Ch <- i
	}
	for _, wk := range workers {
		wk.Stop()
	}
	wg.Wait()
}
