package main

import "fmt"

// Клапан позволяет подключать датчики
type Valve struct {
}

// подключение происходит через интерфейс коннектро
type Connector interface {
	IstallM30Sensor()
}

// По умолчания клапан позволяет подключать датчики M30
func (v *Valve) SensorConnect(connect Connector) {
	fmt.Println("Some sensor connected")
	connect.IstallM30Sensor()
}

// описываем датчик M30
type SensorM30 struct{}

// метод для датчика InstallM30Sensor он отвечает интерфейсу клапана
func (s *SensorM30) IstallM30Sensor() {
	fmt.Println("Sensor M30 connected")
}

// описываем датчика M28
type SensorM28 struct{}

// метод для датчика InstallM28Sensor не отвечает интерфесу клапан, нужен адаптер
func (s *SensorM28) IstallM28Sensor() {
	fmt.Println("Sensor M28 connected")
}

// создаём адаптер, который принимает датчик M28
type SensorM28toM30Adapter struct {
	connect *SensorM28
}

// для адаптера создаём метод InstallM30Sensor, он позволит применять SensorM28 через адаптер
func (s *SensorM28toM30Adapter) IstallM30Sensor() {
	fmt.Println("Sensor M28 connected with adaptor")
	s.connect.IstallM28Sensor()
}

func main() {
	valve := &Valve{}

	// Подключили датчик M30
	sensorM30 := &SensorM30{}
	valve.SensorConnect(sensorM30)

	// Подключаем датчик M28
	sensorM28 := &SensorM28{}
	sensorM28toM30Adapter := &SensorM28toM30Adapter{connect: sensorM28}
	valve.SensorConnect(sensorM28toM30Adapter)
}
