package listener

import (
	"log"
	"net"
	"proxy-server/internal/proxycore"
	"proxy-server/internal/scheduler"
)

func Start(address string, pool *scheduler.WorkerPool, jobQueue *scheduler.JobQueue) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Printf("Listening on %s...", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Parse request in a goroutine
		go func(c net.Conn) {
			req, size, err := proxycore.ParseRequest(c)
			if err != nil {
				c.Close()
				return
			}

			job := scheduler.Job{
				ClientConn: c,
				Request:    req,
				Priority:   size, // placeholder (adjust in Phase 3)
				Size:       size,
			}

			jobQueue.Enqueue(job)
		}(conn)
	}
}
