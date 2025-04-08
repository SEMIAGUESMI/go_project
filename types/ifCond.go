package types

import "fmt"

func varInIf() {
	if i := 10; i < 20 {
		fmt.Println("i is beager than 10")
	}
}

func switchTest(n int) {
	switch n {
	case 1:
		println("monday")
	case 2:
		println("tuesday")
	case 3:
		println("wednesday")
	case 4:
		println("thursday")
	case 5:
		println("friday")
	default:
		println("sunday or saturday")
	}
}

func testSwitch2(day string) {
	switch dayweek := day; dayweek {
	case "monday":
		println("start of week")
	default:
		println("end of week")
	}
}

// Go allows switch without a condition, which acts like an if-else chain.
func testSwitch3(l int) {
	x := l
	switch {
	case x > 10:
		x++
		fmt.Println(x)
		//fallthrough
	case x < 10 && x > 0:
		x--
	default:
		x = 90
	}
	fmt.Println(x)
}
func ControlStructureMain() {
	varInIf()
	switchTest(7)
	testSwitch2("fraiday")
	testSwitch3(60)

}
