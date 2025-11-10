package proxycore

import (
	"io"
	"net"
	"net/http"
	"time"
)

func forwardRequest(clientConn net.Conn, req *http.Request, targetURL string) error {
	// Create new HTTP request for backend server
	newReq, err := http.NewRequest(req.Method, targetURL, req.Body)
	if err != nil {
		return err
	}

	// Copy headers
	for k, v := range req.Header {
		for _, vv := range v {
			newReq.Header.Add(k, vv)
		}
	}

	// Create backend client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Forward request to target server
	resp, err := client.Do(newReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write status line
	_, err = io.WriteString(clientConn, "HTTP/1.1 "+resp.Status+"\r\n")
	if err != nil {
		return err
	}

	// Write headers
	for k, v := range resp.Header {
		for _, vv := range v {
			_, err = io.WriteString(clientConn, k+": "+vv+"\r\n")
			if err != nil {
				return err
			}
		}
	}
	_, err = io.WriteString(clientConn, "\r\n")

	// Write body
	_, err = io.Copy(clientConn, resp.Body)
	return err
}
