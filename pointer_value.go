package main

import (
	"fmt"
)

type Counter struct {
	value int
}

func (c Counter) setValue1(newVal int) {
	fmt.Println("c Counter", c)
	c.value = newVal
	fmt.Println("c.value = newVal", c)
}

func (c *Counter) setValue2(newVal int) {
	fmt.Println(c)
	c.value = newVal
}
func (c Counter) setValue3(newVal int) *Counter {
	c.value = newVal
	return &c
}

func main() {
	//value receiver

	//fmt.Println(counter1)
	//fmt.Println("*counter1.setValue4(20)", (*counter1.setValue4(20)).value)
	//	fmt.Println("counter1", counter1)

	//pointer receiver
	counter2 := Counter{value: 80}
	//counter3 := new(Counter)
	fmt.Println("counter2:=&Counter{value: 80}", counter2)
	counter2.setValue2(30)
	fmt.Println("counter2", counter2)

}
