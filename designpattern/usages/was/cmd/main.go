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

	handler := lib.NewHandler(
		lib.WithStringView(lib.NewStringView()),
		lib.WithJsonView(lib.NewJsonView()),
		lib.WithFileView(lib.NewFileView()),
	)

	lib.NewServlet(handler.Do, listener)
}
