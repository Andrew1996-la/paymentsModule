package smarthome

import (
	"fmt"
	"math/rand"
	"strconv"
)

type MotionSensor struct {
	Id             string
	Name           string
	MotionDetected bool
	Sensitivity    float64
}

func NewMotionSensor(name string) *MotionSensor {
	return &MotionSensor{
		Id:             strconv.Itoa(rand.Int()),
		Name:           name,
		MotionDetected: false,
		Sensitivity:    0,
	}
}

func (m MotionSensor) GetId() string {
	return m.Id
}

func (m MotionSensor) GetName() string {
	return m.Name
}

func (m MotionSensor) GetStatus() string {
	if m.Sensitivity > 0 {
		return fmt.Sprintf("Включено. Уровень чувствительности %.2f", m.Sensitivity)
	}

	return "Выключено"
}

func (m *MotionSensor) SetSensativity(sensativity float64) {
	m.Sensitivity = sensativity
}
