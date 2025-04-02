package main

/*
import "fmt"
import "errors"

// counter
type Counter struct {
	value int
}

func (c Counter) increment() {
	c.value++
}
func (c *Counter) desincrement() {
	c.value--
}

// bank account
type Bankaccount struct {
	balance float32
}

func (account *Bankaccount) deposit(amount float32) {
	account.balance += amount
}
func (account *Bankaccount) withdraw(amount float32) error {
	if account.balance < amount {
		return errors.New("balance insuficient")
	}
	account.balance -= amount
	return nil
}

// rectangle
type Rectangle struct {
	width, height float32
}

func (r Rectangle) area() float32 {
	return r.height * r.width
}
func (r *Rectangle) resize(factor float32) {
	r.height *= factor
	r.width *= factor
}

func main() {
	counter := Counter{value: 5}
	counter.increment()
	fmt.Println(counter.value)
	counter.desincrement()
	fmt.Println(counter.value)

	// bank account

	account := Bankaccount{balance: 1234}
	fmt.Println(account.balance)
	account.deposit(500)
	fmt.Println(account.balance)
	err := account.withdraw(45)
	fmt.Println(account.balance)
	fmt.Println(err)

	// rectangle
	rectangle := Rectangle{5, 9}
	fmt.Println(rectangle.area())
	rectangle.resize(4)
	fmt.Println("height: ", rectangle.height, " width: ", rectangle.height)

}
*/
