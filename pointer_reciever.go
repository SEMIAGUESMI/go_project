package main

import "fmt"

type Counter struct {
	value int
}

func (c Counter) increment() {
	c.value++
}
func (c *Counter) desincrement() {
	c.value--
}

func main() {
	counter := Counter{value: 5}
	counter.increment()
	fmt.Println(counter.value)
	counter.desincrement()
	fmt.Println(counter.value)

}
