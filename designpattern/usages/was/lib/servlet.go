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
		// todo localhost:9999/hello.do 등이 실제 내부로 들어오고 난뒤에 동작하지 않음. 코드 체크 필요!
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
