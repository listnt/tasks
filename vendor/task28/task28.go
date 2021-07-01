package task28

import "fmt"

//Интерфейс класса цели
type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {
}

//Целевой класс, реализет интерфейс
func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

//Адаптируемый класс
type Windows struct{}

//метод необходимый для адаптации
//Как можно заметить, он не реализует интерфейс. Т.е. Классы несовместимы
func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

//Класс адаптера
type WindowsAdapter struct {
	WindowMachine *Windows
}

// Метод адаптирующий метод адаптируемого класса к целевому
// Реализует интерфейс и хранит значение адаптируемого класса
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}

//класс для тестирования
type Client struct {
}

//наш клиент яблочник
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
