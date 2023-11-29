package lib

import (
	"design_patterns/domain"
	"fmt"
	"net"
	"os"
)

type listener struct {
}

func NewListener() domain.Listener {
	return &listener{}
}

func (l listener) Listen(options ...func(config *domain.ListenerConfig)) (net.Listener, error) {
	svr := &domain.ListenerConfig{}
	for _, o := range options {
		o(svr)
	}

	addr, err := net.ResolveTCPAddr(svr.ResolveNetworkType, svr.Address)
	// TODO: return nil, error and decide how to handle it in the calling function
	if err != nil {
		fmt.Println("Failed to resolve address", err.Error())
		os.Exit(1)
	}

	listener, err := net.Listen(svr.NetworkType, addr.String())
	if err != nil {
		fmt.Println("Error listening:", err)
		return nil, err
	}

	return listener, nil
}
