package types

import "fmt"

func Types_declaration() {

	var int_ int = 123
	var bool_ bool
	var float_ float32
	var string_ string
	fmt.Printf(" %v\n %f\n %v\n %q\n",
		int_,
		float_,
		bool_,
		string_)
	fmt.Printf("the typle of the varibale is %T\n ", int_)

	she := 2.4
	fmt.Printf("the typle of the varibale is %T\n", she)

	var1, var2 := 7.009, "ameni is bored"
	fmt.Println(var1, "\n", var2)

	fmt.Println(int(var1))

	const const_ = "nshssn"

}
