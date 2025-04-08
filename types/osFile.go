package types

import (
	"fmt"
	"os"
)

var file *os.File
var err error

func manageOsFile() {
	// 1️⃣ Create or open a file and assign to the global variable
	file, err = os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 2️⃣ Write raw bytes
	_, err = file.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 3️⃣ Write a string
	_, err = file.WriteString(" world ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 5️⃣ Read bytes into a buffer
	buffer := make([]byte, 3)
	n, _ := file.Read(buffer)
	fmt.Println("Content read from file:", string(buffer[:n]))

}
func MainForOsFile() {
	manageOsFile()
}
