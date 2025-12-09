package smarthome

import (
	"math/rand"
	"strconv"
)

type SmartBlob struct {
	Id         string
	Name       string
	IsOn       bool
	Brightness float64 // насколько процентов светит 0-100%
}

func NewSmartBlob(name string) *SmartBlob {
	return &SmartBlob{
		Id:         strconv.Itoa(rand.Int()),
		Name:       name,
		IsOn:       false,
		Brightness: 0,
	}
}

func (s SmartBlob) GetId() string {
	return s.Id
}

func (s SmartBlob) GetName() string {
	return s.Name
}

func (s SmartBlob) GetStatus() string {
	if s.IsOn {
		return "Вклечено"
	}

	return "Выключено"
}

func (s *SmartBlob) TurnOn() error {
	s.IsOn = true
	return nil
}

func (s *SmartBlob) TurnOff() error {
	s.IsOn = false
	return nil
}

func (s SmartBlob) IsOnDevice() bool {
	return s.IsOn
}

func (s *SmartBlob) SetLevel(level float64) {
	s.Brightness = level
}

func (s SmartBlob) GetLevel() float64 {
	return s.Brightness
}
