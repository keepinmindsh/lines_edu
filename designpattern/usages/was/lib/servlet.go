package lib

import (
	"fmt"
	"net"
)

type servlet struct {
}

func NewServlet(handleRequest func(conn net.Conn), listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleRequest(conn)
	}
}
