package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	c  int
}

func (a *Counter) Inc() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.c++
}
func (a *Counter) Get() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.c
}

func main() {
	c := Counter{}
	for i := 0; i < 1000; i++ {
		go c.Inc()
	}

	time.Sleep(time.Second)
	fmt.Println(c.Get())
}
