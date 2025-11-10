package proxycore

import (
	"bufio"
	// "io"
	"log"
	"net"
	"net/http"
	"strings"
)

func HandleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	reader := bufio.NewReader(clientConn)

	// Parse the request from the client
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Printf("Failed to read client request: %v", err)
		return
	}

	targetURL := req.RequestURI
	if !strings.HasPrefix(targetURL, "http") {
		targetURL = "http://" + req.Host + req.RequestURI
	}

	err = forwardRequest(clientConn, req, targetURL)
	if err != nil {
		log.Printf("Error forwarding request: %v", err)
	}
}
