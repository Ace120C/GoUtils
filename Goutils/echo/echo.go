package main

import (
	"fmt"
	"os"
	"strings"
)

var str string

func main()  {
	args := os.Args[1:]
	if len(args) > 0 {
		str = strings.Join(args, " ")
	}
	fmt.Println(str)
}
