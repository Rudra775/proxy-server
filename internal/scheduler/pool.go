package scheduler

import (
	"log"
	"proxy-server/internal/proxycore"
)

type WorkerPool struct {
	workers   int
	queue     *JobQueue
	scheduler SchedulingPolicy
}

func NewWorkerPool(workers int, queue *JobQueue, policy SchedulingPolicy) *WorkerPool {
	return &WorkerPool{
		workers:   workers,
		queue:     queue,
		scheduler: policy,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go func(id int) {
			log.Printf("[Worker %d] Started", id)
			for {
				job := wp.scheduler.PickNext(wp.queue)
				proxycore.ProcessJob(job)
			}
		}(i)
	}
}
