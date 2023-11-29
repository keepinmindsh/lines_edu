package lib

import (
	"design_patterns/domain"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type get struct{}

func (g get) Resolve(param domain.ViewResolverConfig) {
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
	return &get{}
}

type post struct{}

func (p post) Resolve(param domain.ViewResolverConfig) {
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
	return &post{}
}

type fileIo struct{}

func (i fileIo) Resolve(param domain.ViewResolverConfig) {
	fileName, ok := param.Data.(string)
	if ok {
		file, err := os.Open(strings.TrimSpace(fileName)) // For read access.
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close() // make sure to close the file even if we panic.
		n, err := io.Copy(param.Conn, file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n, "bytes sent")
	} else {
		content := "Path is empty"

		fmt.Fprintf(param.Conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(param.Conn, "Content-Length: %d\r\n", len(content))
		fmt.Fprintf(param.Conn, "\r\n")
		fmt.Fprintf(param.Conn, "%s\r\n", content)
	}
}

func NewFileView() domain.ViewResolver {
	return &fileIo{}
}

type httpView struct {
}

func (h httpView) Resolve(parameter domain.ViewResolverConfig) {
	
}

func NewHttpView() domain.ViewResolver {
	return &httpView{}
}