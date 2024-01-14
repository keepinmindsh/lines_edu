package resolver

import (
	"design_patterns/domain"
	"fmt"
)

type stringView struct{}

func (g stringView) Resolve(param domain.ViewResolverConfig) {
	conn := param.Conn
	response, ok := param.Data.(string)
	if ok {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(response))
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, "%s\r\n", response)
	} else {
		response = "Type is not matched!"
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(response))
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, "%s\r\n", response)
	}

}

func NewStringView() domain.ViewResolver {
	return &stringView{}
}
