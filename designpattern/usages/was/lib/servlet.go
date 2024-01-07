package lib

import (
	"design_patterns/domain"
	"net"
)

type Servlet interface {
	Exec(conn net.Conn)
}

type servlet struct {
	ViewResolver map[ResolverType]domain.ViewResolver
}

func (s servlet) Exec(conn net.Conn) {

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
