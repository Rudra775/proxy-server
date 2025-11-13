package scheduler

type SchedulingPolicy interface {
	PickNext(queue *JobQueue) Job
}
