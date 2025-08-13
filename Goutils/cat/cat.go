package main

import (
	"fmt"
	"os"
)

var File string

func main()  {
	Args := os.Args[1]
	if len(os.Args) > 0 {
		File = Args
	}

	f, err := os.Open(File)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	content, err := os.ReadFile(File)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Convert Bytes into a human readable string and outputs it into the console
	os.Stdout.Write(content)

	f.Close()
}
