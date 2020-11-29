package workerpool

import (
	"context"
	"sync"
)

type Task func() error

type WorkerPool struct {
	num     int
	tasks   chan Task
	wg      sync.WaitGroup
	errOnce sync.Once
	err     error
	cancel  context.CancelFunc
}

func (p *WorkerPool) EnqueueTask(i Task) {
	p.tasks <- i
}

func (p *WorkerPool) Wait() error {
	close(p.tasks)
	p.wg.Wait()
	if p.cancel != nil {
		p.cancel()
	}

	return p.err
}

func New(num int, ctx context.Context) (*WorkerPool, context.Context) {
	c, cancel := context.WithCancel(ctx)
	p := &WorkerPool{
		num:    num,
		tasks:  make(chan Task),
		cancel: cancel,
	}

	p.setWorkers()

	return p, c
}

func (p *WorkerPool) setWorkers() {
	for i := 0; i < p.num; i++ {
		p.wg.Add(1)

		go func() {
			defer p.wg.Done()

		Tasks:
			for {
				select {
				case f, ok := <-p.tasks:
					if !ok {
						break Tasks
					}

					if err := f(); err != nil {
						p.errOnce.Do(func() {
							p.err = err

							if p.cancel != nil {
								p.cancel()
							}
						})
					}
				}
			}

		}()
	}
}
