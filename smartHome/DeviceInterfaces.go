package smarthome

// базовый девайс
type Device interface {
	GetId() string
	GetName() string
	GetStatus() string
}

// девайс который можно включать и выключать
type SwitchableDevice interface {
	Device
	TurnOn() error
	TurnOff() error
	IsOnDevice() bool
}

// девайс с датчиками
type Sensor interface {
	Device
	GetValue() float64
}

// девайс с регулировкой
type Adjustable interface {
	Device
	SetLevel(level float64)
	GetLevel() float64
}
