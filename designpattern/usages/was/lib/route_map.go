package lib

import "sync"

var RouteMap *map[string]route

type route struct {
	Url        string
	HttpMethod string
	Exec       interface{}
}

func register(routes *map[string]route) {
	var sc sync.Once

	sc.Do(func() {
		RouteMap = routes
	})
}
