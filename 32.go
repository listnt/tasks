package main

import (
	"fmt"
	"time"
)

func MySleep(d time.Duration) {
	timer1 := time.NewTimer(d)
	<-timer1.C
}

func main() {
	tm := time.Now()
	fmt.Printf("Before sleep, time - %d . %d\n", tm.Second(), tm.Nanosecond())
	MySleep(4 * time.Second)
	tm = time.Now()
	fmt.Printf("After sleep, time - %d . %d\n", tm.Second(), tm.Nanosecond())
}
