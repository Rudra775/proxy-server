package scheduler

type RoundRobinPolicy struct{}

func (p *RoundRobinPolicy) PickNext(queue *JobQueue) Job {
	return queue.Dequeue() // FIFO → round robin
}
