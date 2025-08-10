package main

import (
	"fmt"
	"os"
)

var DirName string

func main()  {
	Args := os.Args[1:]

	for _, DirName := range Args {
		err := os.Mkdir(DirName, 0750)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
