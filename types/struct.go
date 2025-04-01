package types

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Address   Address
}
type Address struct {
	City    string
	Country string
	Street  string
}
type Car_type struct {
	Marke  string
	Model  string
	height int
	width  int
	wheel  struct {
		Radius   int
		Material string
	}
}

type Car struct {
	Make  string
	Model string
}
type truck struct {
	Car
	Bedsize int
}

func (person User) Creat_user() User {
	firstName := person.FirstName
	lastName := person.LastName
	city := person.Address.City
	country := person.Address.Country
	street := person.Address.Street

	return User{firstName, lastName, Address{city, country, street}}

}

func Create_struct() truck {
	lonesTruck := truck{
		Bedsize: 10,
		Car: Car{
			Model: "lamborghini",
			Make:  "e",
		},
	}
	return lonesTruck
}

type Authinfo struct {
	Username string
	Password string
}

func (auth Authinfo) Autorization() string {
	//return "Autorization : Bsic" + auth.Username + ":" + auth.Password
	return fmt.Sprintf("Autorization : Bsic %s:%s",
		auth.Username,
		auth.Password,
	)
}
