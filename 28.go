package main

import "fmt"

//Интерфейс класса цели
type computer interface {
	insertIntoLightningPort()
}

type mac struct {
}

//Целевой класс, реализет интерфейс
func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

//Адаптируемый класс
type windows struct{}

//метод необходимый для адаптации
//Как можно заметить, он не реализует интерфейс. Т.е. Классы несовместимы
func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

//Класс адаптера
type windowsAdapter struct {
	windowMachine *windows
}

// Метод адаптирующий метод адаптируемого класса к целевому
// Реализует интерфейс и хранит значение адаптируемого класса
func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

//класс для тестирования
type client struct {
}

//наш клиент яблочник
func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

func main() {

	client := &client{}
	mac := &mac{}
	//наш клиент доволен, все хорошо. Лайтхнинг к маку
	client.insertLightningConnectorIntoComputer(mac)

	//О нет, тут машина на шиндовс, а клиент яблочник. Что же делать?
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{ // Не беспокойся, программист, ибо я,
		//великий адаптер спешу на помощь

		windowMachine: windowsMachine, //давай сюда свою шиндовс машину
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
