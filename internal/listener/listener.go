package listener

import (
	"log"
	"net"
	"proxy-server/internal/proxycore"
)

func Start(address string) error {
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

		// handle each connection in its own goroutine (concurrency baseline)
		go proxycore.HandleConnection(conn)
	}
}
