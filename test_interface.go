package main

import (
	"LEARN_GO/types"
	"fmt"
)

func main() {
	types.Test(types.BirthdayMessage{
		Name: "ameni",
		Date: "25-08-1997",
	})

	types.Test_computer(types.Apple{Name: "apple", Storage: 64})

	sms := types.Sms{
		Sender_number:  "+393513116571",
		Receipt_number: "+393513116571",
	}
	fmt.Println(types.Test_switch(sms))
}
