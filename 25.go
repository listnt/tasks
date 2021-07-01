package main

import (
	"fmt"
	"sync"
	"time"
)

//структура счетчика
type Counter struct {
	mu sync.Mutex //мютекс синхронизирующий
	c  int        //значение переменной обычное
}

//безопасное инкрементирование
func (a *Counter) Inc() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.c++
}

//безовасное получение значения
func (a *Counter) Get() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.c
}

//код для проверки
func main() {
	c := Counter{}
	for i := 0; i < 1000; i++ {
		go c.Inc()
	}
	time.Sleep(time.Second)
	fmt.Println(c.Get())
}
