package types

type Mystruct struct {
	owner User // owner is private unexported
}

func (o Mystruct) Owner() User {
	return o.owner
}

func (o *Mystruct) SetOwner(newOwner User) {
	o.owner = newOwner
}
