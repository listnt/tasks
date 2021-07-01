package task1

import "fmt"

// Действия осуществляемыми структурами, которые наследуют этот интерфейс
type ActionsIntf interface {
	Action() ActionsIntf /*
		Сущность совершает действие, метод возвращает этот интерфейс
	*/
	Say() ActionsIntf /*
		Сущность говорит, метод возвращает этот интерфейс
	*/
}

type Human struct {
	name   string // Имя человека
	age    int    // Возраст человека
	gender string // Гендер человека
}

type Action struct {
	Human // Композиция человека в структура Action
}

//Base Human Action
func (human *Human) Action() ActionsIntf {
	fmt.Println("I'm going to play World of tanks")
	return human
}

//Base Human Say
func (human *Human) Say() ActionsIntf {
	fmt.Printf("Hello my name is %s. i'm %d years old\n", human.name, human.age)
	return human
}

//New Human
func newHuman(name string, gender string, age int) *Human {
	res := new(Human)
	res.age = age
	res.gender = gender
	res.name = name
	return res
}

//New Action
func NewAction(name string, gender string, age int) *Action {
	res := new(Action)
	res.age = age
	res.gender = gender
	res.name = name
	return res
}

//Child Action Action
func (action *Action) Action() ActionsIntf {
	fmt.Println("I'm going to Eat and Reproduce")
	return &action.Human
}
