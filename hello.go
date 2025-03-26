package main

import (
	"LEARN_GO/types"
	"fmt"
)
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	types.Add_new_data_to_file()
}
