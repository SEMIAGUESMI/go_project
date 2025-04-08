package types

import (
	"fmt"
	"reflect"
)

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

func arrayGo() {
	var arryName [10]int
	arryName[0] = 45
	var arr1 = []int{1, 3, 4}
	fmt.Println(arr1)
	arr := [7]string{"aa", "bb", "cc"}
	fmt.Println(arr)
	arr[2] = "dd"
	arr[3] = "dd" // panic: runtime error: index out of range [3] with length 3 for arr := []string{"aa", "bb", "cc"}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, val := range arr {
		fmt.Println("index : ", index, " : ", val)
	}
	fmt.Println(reflect.TypeOf(arr1))
}

func slice() {
	var s1 []int
	fmt.Println(s1)
	s1 = append(s1, 1, 2, 3, 4, 5, 7)
	fmt.Println(s1)

	var s2 = []string{"dd", "ll", "ii"}
	fmt.Println(s2)
	s3 := []float32{8.0, 9.6}
	fmt.Println(s3)
	arr1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	//create slice by copying from array
	s4 := arr1[1:6]
	fmt.Println(s4)

	for _, value := range s4 {
		s4 = append(s4, value*2)
	}
	fmt.Println(s4)
}

func sliceWithMake() {
	s1 := make([]int, 3)
	s1[0], s1[1], s1[2] = 1, 2, 3
	fmt.Println(s1)

	s2 := make([]int, 0)
	for i := 0; i < 5; i++ {
		s2 = append(s2, i*3)
	}
	fmt.Println(s2)

	s3 := s2[:3]
	fmt.Println(s3)

	for index, value := range s3 {
		fmt.Println("range", index, value)
	}

}

func MainTypes() {
	sliceWithMake()
}
