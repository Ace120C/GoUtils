package main

import (
	"fmt"
	"os"
)

func main()  {
	// getdw gets the current location
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(wd)
}
