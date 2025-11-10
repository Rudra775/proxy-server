package scheduler

import (
	"net"
	"net/http"
)

type Job struct {
	ClientConn net.Conn
	Request    *http.Request
	Priority   int // used for priority scheduling
	Size       int // used for SJF scheduling (approx bytes)
}
