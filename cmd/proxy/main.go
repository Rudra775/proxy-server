package main

import (
	"log"
	"proxy-server/internal/listener"
)

func main() {
	addr := ":8080"
	log.Printf("Starting proxy server on %s\n", addr)

	err := listener.Start(addr)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}
}
