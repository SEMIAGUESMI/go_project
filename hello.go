package main

import (
	"LEARN_GO/types"
	"fmt"
)
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	//types.Condistional(-1)
	//types.Base()
	_, name := types.Name()
	fmt.Println("my name is ", name)

}
