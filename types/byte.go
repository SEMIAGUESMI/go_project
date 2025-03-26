package types

import (
	"fmt"
	"os"
)

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
func Write_file_using_byte() {
	data := "ameni is learning golang"
	dat_to_byte := []byte(data)
	err := os.WriteFile("file", dat_to_byte, 0644)
	if err != nil {
		fmt.Println("error in writing file", err)
		return
	}
	fmt.Println("successufly write")
}

func Read_file_into_byte() {
	data, err := os.ReadFile("./types/file")
	if err != nil {
		fmt.Println("error in writing file : ", err)
		return
	}
	fmt.Println(data)
	fmt.Println(string(data))

}
