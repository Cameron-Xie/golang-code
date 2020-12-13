package elasticwp

import (
	"context"
	"sync"
)

type Task func() error

type WorkerPool struct {
	max     int
	wg      sync.WaitGroup
	tasks   chan Task
	workers chan struct{}
	cancel  context.CancelFunc
	errOnce sync.Once
	err     error
}

func (p *WorkerPool) EnqueueTask(i Task) {
	p.tasks <- i
}

func (p *WorkerPool) Wait() error {
	close(p.tasks)
	p.wg.Wait()
	return p.err
}

func New(max int, ctx context.Context) (*WorkerPool, context.Context) {
	c, cancel := context.WithCancel(ctx)
	p := &WorkerPool{
		max:     max,
		tasks:   make(chan Task),
		workers: make(chan struct{}),
		cancel:  cancel,
	}

	go p.setupWorkers()

	return p, c
}

func (p *WorkerPool) setupWorkers() {
	defer close(p.workers)
	workerNum := 0

	for task := range p.tasks {
		if workerNum < p.max {
			go p.registerNewWorker()
			workerNum++
		}

		<-p.workers
		p.wg.Add(1)
		go p.processTask(task)
	}

	for workerNum > 0 {
		<-p.workers
		workerNum--
	}
}

func (p *WorkerPool) registerNewWorker() {
	p.workers <- struct{}{}
}

func (p *WorkerPool) processTask(task Task) {
	defer p.wg.Done()

	if err := task(); err != nil {
		p.errOnce.Do(func() {
			p.err = err

			if p.cancel != nil {
				p.cancel()
			}
		})
	}

	p.registerNewWorker()
}
