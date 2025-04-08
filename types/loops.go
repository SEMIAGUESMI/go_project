package types

import "fmt"

func basicFor() {
	for i := 0; i < 10000; i++ {
		fmt.Println("saif twa7achtek !")
	}
}
func forAsWhile() {
	i := 0
	for i <= 5 {
		i++
		fmt.Println(i)
	}
}

func infinitLoop() {
	for {
		fmt.Println("hello guys!")
	}
}

var numbers = []int{10, 23, 90}

func rangeLoop() {
	for index, value := range numbers {
		fmt.Println(index, value)
	}
}

func rangeLoopString(s string) {
	for pos, c := range s {
		fmt.Printf("%v ekrfjre %v", pos, c)
	}
}
func MainLoops() {
	rangeLoopString("ameni guesmi")
}
