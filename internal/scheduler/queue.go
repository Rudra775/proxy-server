package scheduler

type JobQueue struct {
	jobs chan Job
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		jobs: make(chan Job, size),
	}
}

func (q *JobQueue) Enqueue(job Job) {
	q.jobs <- job
}

func (q *JobQueue) Dequeue() Job {
	return <-q.jobs
}

func (q *JobQueue) Length() int {
	return len(q.jobs)
}
