package main

import (
	"design_patterns/domain"
	"design_patterns/lib"
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

	servlet := lib.NewServlet(lib.SettingResolver(map[lib.ResolverType]domain.ViewResolver{
		lib.JsonView:      lib.NewJsonView(),
		lib.StringView:    lib.NewStringView(),
		lib.MultiPartView: lib.NewMultiPartView(),
		lib.HttpView:      lib.NewHttpView(),
	}))

	lib.NewServer(servlet.Exec, listener)
}
