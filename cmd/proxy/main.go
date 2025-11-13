package main

import (
	"log"
	"proxy-server/internal/listener"
	"proxy-server/internal/scheduler"
)

func main() {
	addr := ":8080"
	jobQueue := scheduler.NewJobQueue(100)
	policy := &scheduler.RoundRobinPolicy{}
	workerPool := scheduler.NewWorkerPool(10, jobQueue, policy)
	workerPool.Start()

	log.Printf("Starting proxy server on %s\n", addr)

	err := listener.Start(addr, workerPool, jobQueue)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}
}
