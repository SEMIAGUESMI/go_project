package types

import "fmt"

func Byte_8() {
	var a byte = 65
	fmt.Println(a)
	fmt.Println(string(a))
}
func Byte_string() {
	str := "ameni"
	byte_str := []byte(str) // Convert string to byte slice
	fmt.Println(byte_str)
	byte_str[4] = 'M'
	fmt.Println(byte_str)
	fmt.Println(string(byte_str))

}
