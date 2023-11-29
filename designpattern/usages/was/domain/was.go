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

type ViewResolverConfig struct {
	Conn net.Conn
	Data interface{}
}

type ViewResolver interface {
	Resolve(parameter ViewResolverConfig)
}
