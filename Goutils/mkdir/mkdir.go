package main

import (
	"fmt"
	"os"
	"strings"
)

var DirName string

func main()  {
	Args := os.Args[1:]
	if len(os.Args) > 0 {
		DirName = strings.Join(Args, " ")
	}

	for _, DirName := range Args {
		err := os.Mkdir(DirName, 0750)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
