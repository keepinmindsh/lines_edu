package servlet

import (
	"design_patterns/domain"
	"net"
	"net/http"
)

func (s servletContainer) Do(method string, path string, request string, conn net.Conn) {
	if method == "GET" {
		if path == "/hello" {
			s.ViewResolver[domain.StringView].Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: "Hello World!",
			})
		} else if path == "/hello.do" {
			s.ViewResolver[domain.HttpView].Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: nil,
			})
		} else if path == "/hello.png" {
			s.ViewResolver[domain.MultiPartView].Resolve(domain.ViewResolverConfig{
				Conn: conn,
				Data: "hello.png",
			})
		} else {
			writeErrorResponse(conn, http.StatusNotFound)
		}
	} else if method == "POST" {
		s.ViewResolver[domain.JsonView].Resolve(domain.ViewResolverConfig{
			Conn: conn,
			Data: request,
		})
	} else if method == "PUT" {
		s.ViewResolver[domain.JsonView].Resolve(domain.ViewResolverConfig{
			Conn: conn,
			Data: request,
		})
	} else if method == "DELETE" {
		s.ViewResolver[domain.JsonView].Resolve(domain.ViewResolverConfig{
			Conn: conn,
			Data: request,
		})
	} else {
		writeErrorResponse(conn, http.StatusMethodNotAllowed)
	}

}
