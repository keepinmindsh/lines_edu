package main

type (
	Domain interface {
		Hi()
		Hello()
		Bye()
	}
)

type Service struct {
}

func (s Service) Hi() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Hello() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Bye() {
	//TODO implement me
	panic("implement me")
}

func NewService() Domain {
	return Service{}
}

type subscriber struct {
	domain Domain
}

func Subscriber(domain Domain) {
	_ = subscriber{
		domain: domain,
	}
}

func main() {
	Subscriber(NewService())
}
