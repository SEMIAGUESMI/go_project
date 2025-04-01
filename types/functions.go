package types

import "fmt"

func Conditional(a int) {
	if a == 0 {
		a++
	} else if a > 0 {
		a += 2
	} else {
		a = 20
	}
	fmt.Println(a)

}
func Concat(str1 string, str2 string) string {
	var1 := 7.009
	var1 = 0.9
	fmt.Println(var1)
	return str1 + str2
}
func Base() {
	age := 20
	fmt.Println(age)
	age = increment(age, 7)
	fmt.Println(age)
}
func increment(initial_age int, years int) int {

	return initial_age + years
}
func Name() (string, string) {
	return "semia", "guesmi"
}
func NAKED_Return() (x, y int, name string) {
	x = 1
	y = 2
	name = "ameni"
	return
}
