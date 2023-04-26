package main

type TheWayToGoBuSan interface {
	UseTransportationInSeoul()
	Moving()
	ArriveToDestination()
}

type UseBus struct{}

func NewTheWayBus() TheWayToGoBuSan {
	return &UseBus{}
}

func (ub *UseBus) UseTransportationInSeoul() {

}

func (ub *UseBus) Moving() {

}

func (ub *UseBus) ArriveToDestination() {

}

func main() {
	howToGoBusan := NewTheWayBus()

	howToGoBusan.UseTransportationInSeoul()

	howToGoBusan.Moving()

	howToGoBusan.ArriveToDestination()
}
