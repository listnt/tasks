package main

import (
	"task1"
)

//driver code
func main() {
	var Act task1.ActionsIntf
	Act = task1.NewAction("Igor", "NA", 32)
	Act.Action().Say().Action()
}
