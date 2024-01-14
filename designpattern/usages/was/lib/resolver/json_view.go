package resolver

import (
	"design_patterns/domain"
	"fmt"
	"strings"
)

type jsonView struct{}

func (p jsonView) Resolve(param domain.ViewResolverConfig) {
	conn := param.Conn
	request, ok := param.Data.(string)

	if ok {
		content := strings.Split(request, "\r\n\r\n")[1]

		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(content))
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, "%s\r\n", content)
	} else {
		content := "Content is empty"

		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(content))
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, "%s\r\n", content)
	}
}

func NewJsonView() domain.ViewResolver {
	return &jsonView{}
}
