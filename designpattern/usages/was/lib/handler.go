package lib

import (
	"design_patterns/domain"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type handler struct {
	StringViewResolver domain.ViewResolver
	JsonViewResolver   domain.ViewResolver
	FileViewResolver   domain.ViewResolver
}

func NewHandler(options ...func(config *handler)) domain.ServletHandler {
	svr := &handler{}
	for _, o := range options {
		o(svr)
	}

	return &handler{
		StringViewResolver: svr.StringViewResolver,
		JsonViewResolver:   svr.JsonViewResolver,
		FileViewResolver:   svr.FileViewResolver,
	}
}

func WithStringView(get domain.ViewResolver) func(*handler) {
	return func(h *handler) {
		h.StringViewResolver = get
	}
}

func WithJsonView(post domain.ViewResolver) func(*handler) {
	return func(h *handler) {
		h.JsonViewResolver = post
	}
}

func WithFileView(io domain.ViewResolver) func(*handler) {
	return func(h *handler) {
		h.FileViewResolver = io
	}
}

func (h handler) Do(conn net.Conn) {
	buf := make([]byte, 1024)

	// Read the request header
	n, err := conn.Read(buf)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return
		} else {
			fmt.Println(err.Error())
		}
	}

	request := string(buf[:n])

	// Parse the request method and path
	method, path := parseRequestLine(request)

	// Resolve the GET method
	if method == "GET" {
		if path == "/hello" {
			h.StringViewResolver.Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: "Hello World!",
			})
		} else if path == "/hello.do" {

		} else if path == "/hello.png" {
			h.FileViewResolver.Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: "hello.png",
			})
		} else {
			writeErrorResponse(conn, http.StatusNotFound)
		}
	} else if method == "POST" {
		h.JsonViewResolver.Resolve(domain.ViewResolverConfig{
			Conn: conn,
			Data: request,
		})
	} else {
		writeErrorResponse(conn, http.StatusMethodNotAllowed)
	}

	conn.Close()
}

func parseRequestLine(request string) (string, string) {
	parts := strings.Split(request, "\r\n")

	var split []string
	for _, part := range parts {
		if strings.Contains(part, "HTTP") {
			split = strings.Split(parts[0], " ")
			break
		}
	}

	return split[0], split[1]
}

func writeErrorResponse(conn net.Conn, status int) {
	fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", status, http.StatusText(status))
	fmt.Fprintf(conn, "Content-Length: 0\r\n")
	fmt.Fprintf(conn, "\r\n")
}
