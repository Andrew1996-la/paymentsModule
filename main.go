package main

import (
	"fmt"
	smarthome "study/smartHome"

	"github.com/k0kubun/pp"
)

func getStatusesDevices(devices []smarthome.Device) {
	pp.Println("Провека работоспособности устройств")
	for _, device := range devices {
		pp.Printf("Девайс: %s\nСтатус: %s\n\n", device.GetName(), device.GetStatus())
	}
}

func main() {
	smartBlobLivingroom := smarthome.NewSmartBlob("свет в зале")
	moutionSensorInGarden := smarthome.NewMotionSensor("датчик движения во дворе")
	smartTVSocket := smarthome.NewSmartSocket("умная TV розетка")

	devices := []smarthome.Device{smartBlobLivingroom, moutionSensorInGarden, smartTVSocket}

	getStatusesDevices(devices)

	pp.Println("Включаю swithcable устройства")
	switchables := []smarthome.SwitchableDevice{smartBlobLivingroom, smartTVSocket}
	for _, device := range switchables {
		device.TurnOn()
		pp.Printf("%s : %s\n", device.GetName(), device.GetStatus())
	}

	fmt.Println("")

	pp.Println("Влючаю sensor устройство")
	moutionSensorInGarden.SetSensativity(7)
	pp.Println(moutionSensorInGarden.GetStatus())

	fmt.Println("")

	getStatusesDevices(devices)
}
