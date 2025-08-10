package main

import (
	"fmt"
	"os"
	"strings"
)

var FileName string

func main()  {
	Args := os.Args[1:]
	if len(os.Args) > 0 {
		FileName = strings.Join(Args, " ")
	}

	f, err := os.Create(FileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	f.Close()
}
