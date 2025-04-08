package types

import (
	"errors"
	"fmt"
	"math"
)

type shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
	Width  float32
	Lenght float32
}
type cercle struct {
	radius float32
}

func (r Rectangle) Area() float32 {

	return r.Lenght * r.Width
}
func (r Rectangle) Perimeter() float32 {
	return 2 * r.Width * r.Lenght
}

func (c cercle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c cercle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

type Message interface {
	GetMessage() string
}

type BirthdayMessage struct {
	Name string
	Date string
}
type ReportMessage struct {
	Title string
	Date  string
}

func Test(m Message) string {
	return m.GetMessage()
}

func (bm BirthdayMessage) GetMessage() string {
	return "birthday of " + bm.Name + "on" + bm.Date
}

func (bm ReportMessage) GetMessage() string {
	return "birthday of " + bm.Title + "on" + bm.Date
}

type Computer interface {
	getinfo() (string, int)
}
type Apple struct {
	Name    string
	Storage int
}
type Hp struct {
	Name    string
	Storage int
}

func (appl Apple) getinfo() (string, int) {
	return appl.Name, appl.Storage
}
func (appl Hp) getinfo() (string, int) {
	return appl.Name, appl.Storage
}
func TestComputer(computer Computer) {
	computer.getinfo()
}

type message interface {
	getSender() string
}
type Email struct {
	SenderAddress  string
	ReceiptAddress string
}
type Sms struct {
	SenderNumber  string
	ReceiptNumber string
}

func (e Email) getSender() string {
	return e.SenderAddress
}

func (s Sms) getSender() string {
	return s.SenderNumber
}

func Test_message(m message) string {
	em, ok := m.(Email)
	if ok {
		return em.getSender()
	}
	c, ok := m.(Sms)
	if ok {
		return c.getSender()
	}
	return ""

}
func TestSwitch(m message) string {
	switch v := m.(type) {
	case Email:
		return v.getSender()
	case Sms:
		return v.getSender()
	default:
		return ""
	}

}

// empty interface
var I interface{} = "hello world"

func EmptyiInterface() {
	s, ok := I.(int)
	fmt.Println(ok)
	if ok {
		fmt.Println(s)
	}
}

func SwitchInterface(i interface{}) error {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	default:
		return errors.New("type not available")
	}
	return nil
}

func InterfaceMain() {
	address := Address{City: "Milan", Country: "Italy", Street: "via poscolle"}
	user := User{FirstName: "ameni", LastName: "guesmi", Address: address}

	obj := Mystruct{}
	fmt.Println(obj)

	obj.SetOwner(user)
	fmt.Println(obj)

	fmt.Println(I)
	EmptyiInterface()

	var k interface{} = 2222.0
	k = "ekdjwjdew"
	err := SwitchInterface(k)
	if err != nil {
		fmt.Println(err)
	}
}
