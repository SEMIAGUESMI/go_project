package types

import "math"

// create method for struct
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
func Test_computer(computer Computer) {
	computer.getinfo()
}
