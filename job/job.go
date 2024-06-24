package job

import (
	"sync"
)

type Job func()

type Pool struct {
	JobQueue chan Job
	wg       sync.WaitGroup
}

var closeOnce sync.Once

// NewPool creates a new pool with multiple workers
func NewPool(maxWorkers *int) *Pool {
	p := Pool{JobQueue: make(chan Job)}
	p.wg.Add(*maxWorkers)

	for i := 0; i < *maxWorkers; i++ {
		go func() {
			defer p.wg.Done()
			for job := range p.JobQueue {
				job()
			}
		}()
	}
	return &p
}

// AddJobTar adds a job to the pool worker queue
func (p *Pool) AddJobTar(job func(s, t string), arg1, arg2 string) {
	p.JobQueue <- func() { job(arg1, arg2) }
}

// AddJobUnTar adds a job to the pool worker queue
func (p *Pool) AddJobUnTar(job func(s string), arg string) {
	p.JobQueue <- func() { job(arg) }
}

// Wait Wait to complete
func (p *Pool) Wait() {
	closeOnce.Do(func() {
		close(p.JobQueue)
	})
	p.wg.Wait()
}
