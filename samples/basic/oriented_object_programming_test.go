package main

import "testing"

// TODO : 만약 고속버스가 아닌 비행기를 이용해야 한다면?
// TODO : 만약 걸어가야 한다면?

func Test_Main(t *testing.T) {

	factory := Factory(Bus{})

	factory.FromMyHome()
	factory.GoToPlatform()
	factory.GetOnTheTransportation()
	factory.ArrivedAtDestination()
}

func Factory(typeValue interface{}) Transportation {
	switch typeValue {
	case typeValue.(Bus):
		return &Bus{}
	default:
		return nil
	}
}

type (
	Transportation interface {
		FromMyHome()
		GetOnTheTransportation()
		GoToPlatform()
		ArrivedAtDestination()
	}
)

type Airplane struct {
}

type Bus struct {
}

func (b Bus) FromMyHome() {
	//TODO implement me
	panic("implement me")
}

func (b Bus) GetOnTheTransportation() {
	//TODO implement me
	panic("implement me")
}

func (b Bus) GoToPlatform() {
	//TODO implement me
	panic("implement me")
}

func (b Bus) ArrivedAtDestination() {
	//TODO implement me
	panic("implement me")
}

func NewTransportation() Transportation {
	return &Bus{}
}
