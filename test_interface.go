package main

import "LEARN_GO/types"

func main() {
	types.Test(types.BirthdayMessage{
		Name: "ameni",
		Date: "25-08-1997",
	})

	types.Test_computer(types.Apple{Name: "apple", Storage: 64})

}
