package main

import (
	"fmt"
	"os"
)

var str string = "y"

func main()  {
	// os.Args provides access to raw command-line arguments
	if len(os.Args) > 1 {
		str = os.Args[1]
	}

	for true {
		fmt.Println(str)
	}
}
