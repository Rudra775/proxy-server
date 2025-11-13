package proxycore

import (
	"/internal/scheduler"
	"log"
)

func ProcessJob(job scheduler.Job) {
	err := forwardRequest(job.ClientConn, job.Request, job.Request.URL.String())
	if err != nil {
		log.Printf("Error processing job: %v", err)
		job.ClientConn.Close()
	}
}
