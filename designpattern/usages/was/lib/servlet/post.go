package servlet

import (
	"fmt"
	"net"
	"net/http"
)

func (s servletContainer) Post(conn net.Conn) {
	writeErrorResponse(conn, http.StatusMethodNotAllowed)
}

func writeErrorResponse(conn net.Conn, status int) {
	fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", status, http.StatusText(status))
	fmt.Fprintf(conn, "Content-Length: 0\r\n")
	fmt.Fprintf(conn, "\r\n")
}
