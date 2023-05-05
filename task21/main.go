package main

import "fmt"

type (
	Headphones struct {
		volume int
	}
	//оборачивает Headphones, чтобы в реализации интерфейса обращаться к методу Headphones
	HeadphonesAdapter struct {
		*Headphones
	}
	//чтобы объект мог подключиться по bluetooth, он должен иметь метод
	Bluetooth interface {
		BluetoothConnect()
	}
)

// сами наушники могут только издавать звук
func (headphones *Headphones) MakeSound(sound string) string {
	return fmt.Sprintf("%s (with volume %v db)", sound, headphones.volume)
}

// поэтому нам нужна структура-адаптер, которая умеет подключаться по BT, а после подключения заставляет наушники издать звук подключения
func (adapter *HeadphonesAdapter) BluetoothConnect() {
	fmt.Printf("Наушники подключены по Bluetooth\n")
	fmt.Println(adapter.Headphones.MakeSound("connected"))
}
func main() {
	var blueTooth Bluetooth
	blueTooth = &HeadphonesAdapter{
		Headphones: &Headphones{
			volume: 20,
		},
	}
	blueTooth.BluetoothConnect()
}
