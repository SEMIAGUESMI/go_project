package types

import "fmt"

func Test_Rune() {
	var k rune = '🚀'
	r := "GO🚀" // r with type rune
	fmt.Println(k)
	fmt.Println(string(k))
	for _, unicode := range r {
		fmt.Printf("unicode caracter is %c\n", unicode)
	}

	for i, unicode := range r {
		fmt.Printf("caracter index %d : %c \n", i, unicode)
	}
}
