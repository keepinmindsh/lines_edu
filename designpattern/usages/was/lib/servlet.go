package lib

import (
	"design_patterns/domain"
	"design_patterns/lib/servlet"
	"net"
)

type Servlet interface {
	Exec(conn net.Conn)
}

type servletInstance struct {
	servlet domain.Servlet
}

func (s servletInstance) Exec(conn net.Conn) {
	pre, err := s.servlet.Pre(conn)

	if err == nil {
		s.servlet.Do(pre.Method, pre.Path, "", conn)
	}

	if err != nil {
		s.servlet.Post(conn)
	}

	conn.Close()
}

func NewServletInstance(viewResolver map[domain.ResolverType]domain.ViewResolver) Servlet {
	return &servletInstance{
		servlet: servlet.NewServletContainer(servlet.SettingResolver(viewResolver)),
	}
}
