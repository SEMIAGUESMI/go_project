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
	err := os.WriteFile("./types/file.txt", dat_to_byte, 0634)
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
func Add_new_data_to_file() {
	file, err := os.OpenFile("./types/file.txt", os.O_APPEND|os.O_WRONLY, 6453)
	if err != nil {
		fmt.Println("error in reading file : ", err)
		return
	}
	defer file.Close() // Ensure the file is closed after writing

	newdata := []byte(" and she is going on quiqly !")

	_, err = file.Write(newdata)
	if err != nil {
		fmt.Println("error in writing file : ", err)
		return
	}
	fmt.Println("Data appended successfully!")
}
