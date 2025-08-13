package main

import (
	"fmt"
	"os"
	"flag"
)


func main()  {
	force := flag.Bool("f", false, "Remove content forcefully")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No File Specified")
		return
	}

	// for now it only checks for one file/directory at a time
	File := flag.Args()[0]
	// checks if the files exist and returns a description of said file
	fi, err := os.Stat(File)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// checking if this is a directory
	if fi.IsDir() {
		if *force {
			err = os.RemoveAll(File)
			if err != nil {
				fmt.Println("Error: ", err)
			} 
		} 			
	} else {
		err = os.Remove(File)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
