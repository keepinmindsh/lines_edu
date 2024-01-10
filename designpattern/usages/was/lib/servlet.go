package lib

import (
	"design_patterns/domain"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type Servlet interface {
	Exec(conn net.Conn)
}

type servlet struct {
	ViewResolver map[ResolverType]domain.ViewResolver
}

type parseResult struct {
	Err         error
	Method      string
	Path        string
	ContentType string
}

func (s servlet) Exec(conn net.Conn) {
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
	parseResult := parseRequestLine(request)

	// todo content 타입을 체크하는 로직
	if parseResult.Err != nil {
		writeErrorResponse(conn, http.StatusNoContent)
		conn.Close()
		return
	}

	// Resolve the GET method
	if parseResult.Method == "GET" {
		if parseResult.Path == "/hello" {
			s.ViewResolver[StringView].Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: "Hello World!",
			})
		} else if parseResult.Path == "/hello.do" {
			s.ViewResolver[HttpView].Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: nil,
			})
		} else if parseResult.Path == "/hello.png" {
			s.ViewResolver[MultiPartView].Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: "hello.png",
			})
		} else {
			writeErrorResponse(conn, http.StatusNotFound)
		}
	} else if parseResult.Method == "POST" {
		s.ViewResolver[JsonView].Resolve(domain.ViewResolverConfig{
			Conn: conn,
			Data: request,
		})
	} else if parseResult.Method == "PUT" {
		s.ViewResolver[JsonView].Resolve(domain.ViewResolverConfig{
			Conn: conn,
			Data: request,
		})
	} else if parseResult.Method == "DELETE" {
		s.ViewResolver[JsonView].Resolve(domain.ViewResolverConfig{
			Conn: conn,
			Data: request,
		})
	} else {
		writeErrorResponse(conn, http.StatusMethodNotAllowed)
	}

	conn.Close()
}

func parseRequestLine(request string) parseResult {
	log.Print(request)

	parts := strings.Split(request, "\r\n")

	var split []string
	var contentType string
	for _, part := range parts {
		if strings.Contains(part, "HTTP") {
			split = strings.Split(parts[0], " ")
			break
		}

		if strings.Contains(part, "Content-Type") {
			contentType = strings.Split(part, " ")[1]
		}
	}

	if contentType == "" {
		return parseResult{
			Err:         errors.New("content-Type is empty, please check your header info"),
			Method:      "",
			Path:        "",
			ContentType: "",
		}
	}

	return parseResult{
		Err:         nil,
		Method:      split[0],
		Path:        split[1],
		ContentType: contentType,
	}
}

func writeErrorResponse(conn net.Conn, status int) {
	fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", status, http.StatusText(status))
	fmt.Fprintf(conn, "Content-Length: 0\r\n")
	fmt.Fprintf(conn, "\r\n")
}

func NewServlet(options ...func(config *servlet)) Servlet {
	svr := &servlet{}
	for _, o := range options {
		o(svr)
	}

	return &servlet{
		ViewResolver: svr.ViewResolver,
	}
}

func SettingResolver(resolver map[ResolverType]domain.ViewResolver) func(*servlet) {
	return func(h *servlet) {
		h.ViewResolver = resolver
	}
}
