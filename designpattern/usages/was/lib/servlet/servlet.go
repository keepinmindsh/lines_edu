package servlet

import "design_patterns/domain"

type servletContainer struct {
	ViewResolver map[domain.ResolverType]domain.ViewResolver
}

func NewServletContainer(options ...func(config *servletContainer)) domain.Servlet {
	svr := &servletContainer{}
	for _, o := range options {
		o(svr)
	}

	return svr
}

func SettingResolver(resolver map[domain.ResolverType]domain.ViewResolver) func(*servletContainer) {
	return func(h *servletContainer) {
		h.ViewResolver = resolver
	}
}
