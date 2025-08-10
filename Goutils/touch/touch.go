package main

import (
	"fmt"
	"os"
)

var FileName string

func main()  {
	Args := os.Args[1:]

	for _, FileName := range Args {
		f, err := os.Create(FileName)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		defer f.Close()
	}

}
