package proxycore

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

func ParseRequest(clientConn net.Conn) (*http.Request, int, error) {
	reader := bufio.NewReader(clientConn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Printf("Parse error: %v", err)
		return nil, 0, err
	}

	size := int(req.ContentLength)
	return req, size, nil
}
