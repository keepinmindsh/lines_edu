package lib

type handleMap struct {
	URL        string
	HttpMethod string
	Data       interface{}
}

type handlerMapping interface {
	Do(args handleMap)
}