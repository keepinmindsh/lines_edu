package main

import (
	"design_patterns/domain"
	"design_patterns/lib"
	"design_patterns/lib/resolver"
)

func main() {
	listenClient := lib.NewListener()

	listener, err := listenClient.Listen(
		domain.WithAddress(":9999"),
		domain.WithNetworkType("tcp"),
		domain.WithResolveNetworkType("tcp4"),
	)
	if err != nil {
		panic(err)
	}

	servlet := lib.NewServletInstance(map[domain.ResolverType]domain.ViewResolver{
		domain.JsonView:      resolver.NewJsonView(),
		domain.StringView:    resolver.NewStringView(),
		domain.MultiPartView: resolver.NewMultiPartView(),
		domain.HttpView:      resolver.NewHttpView(),
	})

	lib.NewServer(servlet.Exec, listener)
}
