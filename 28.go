package main

import "task28"

func main() {
	client := &task28.Client{}
	mac := &task28.Mac{}
	//наш клиент доволен, все хорошо. Лайтхнинг к маку
	client.InsertLightningConnectorIntoComputer(mac)

	//О нет, тут машина на шиндовс, а клиент яблочник. Что же делать?
	windowsMachine := &task28.Windows{}
	windowsMachineAdapter := &task28.WindowsAdapter{ // Не беспокойся, программист, ибо я,
		//великий адаптер спешу на помощь

		WindowMachine: windowsMachine, //давай сюда свою шиндовс машину
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
