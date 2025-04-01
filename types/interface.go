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

type message interface {
	get_sender() string
}
type Email struct {
	Sender_address  string
	Receipt_address string
}
type Sms struct {
	Sender_number  string
	Receipt_number string
}

func (e Email) get_sender() string {
	return e.Sender_address
}
func (s Sms) get_sender() string {
	return s.Sender_number
}

func Test_message(m message) string {
	em, ok := m.(Email)
	if ok {
		return em.get_sender()
	}
	c, ok := m.(Sms)
	if ok {
		return c.get_sender()
	}
	return ""

}
func Test_switch(m message) string {
	switch v := m.(type) {
	case Email:
		return v.get_sender()
	case Sms:
		return v.get_sender()
	default:
		return ""
	}

}
