package domain

import "net"

type ListenerConfig struct {
	ResolveNetworkType string
	NetworkType        string
	Address            string
}

type Listener interface {
	Listen(options ...func(config *ListenerConfig)) (net.Listener, error)
}

func WithResolveNetworkType(resolveNetworkType string) func(*ListenerConfig) {
	return func(config *ListenerConfig) {
		config.ResolveNetworkType = resolveNetworkType
	}
}

func WithNetworkType(networkType string) func(*ListenerConfig) {
	return func(config *ListenerConfig) {
		config.NetworkType = networkType
	}
}

func WithAddress(address string) func(*ListenerConfig) {
	return func(config *ListenerConfig) {
		config.Address = address
	}
}

type ServletHandler interface {
	Do(conn net.Conn)
}

type ResolverType string

const (
	StringView    ResolverType = "StringViewResolver"
	JsonView      ResolverType = "JsonViewResolver"
	MultiPartView ResolverType = "MultiPartView"
	HttpView      ResolverType = "HttpView"
)

type ViewResolverConfig struct {
	Conn net.Conn
	Data interface{}
}

type ViewResolver interface {
	Resolve(parameter ViewResolverConfig)
}

type ParseResult struct {
	Err         error
	Method      string
	Path        string
	ContentType string
}

type Servlet interface {
	Pre(conn net.Conn) (ParseResult, error)
	Do(method, path, request string, conn net.Conn)
	Post(conn net.Conn)
}
