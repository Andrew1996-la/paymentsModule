package smarthome

import (
	"math/rand"
	"strconv"
)

type SmartSocket struct {
	Id   string
	Name string
	IsOn bool
}

func NewSmartSocket(name string) *SmartSocket {
	return &SmartSocket{
		Id:   strconv.Itoa(rand.Int()),
		Name: name,
		IsOn: false,
	}
}

func (s SmartSocket) GetId() string {
	return s.Id
}

func (s SmartSocket) GetName() string {
	return s.Name
}

func (s SmartSocket) GetStatus() string {
	if s.IsOn {
		return "Вклечено"
	}

	return "Выключено"
}

func (s *SmartSocket) TurnOn() error {
	s.IsOn = true
	return nil
}

func (s *SmartSocket) TurnOff() error {
	s.IsOn = false
	return nil
}

func (s SmartSocket) IsOnDevice() bool {
	return s.IsOn
}
