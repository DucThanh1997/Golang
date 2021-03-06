package workerpool2

import (
	"fmt"
	"sync"
)

type Pool struct {
	Tasks   		[]*Task

	concurrency   	int
	collector     	chan *Task
	wg            	sync.WaitGroup
}


func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		collector:   make(chan *Task, 1000),
	}
}
// Run runs all work within the pool and blocks until it's
// finished.
func (p *Pool) Run() {
	for i := 1; i <= p.concurrency; i++ {
		fmt.Println("i")
		worker := NewWorker(p.collector, i)
		worker.Start(&p.wg)
		
	}

	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}
	close(p.collector)

	p.wg.Wait()
}