package main

import (
	"fmt"
	"os/user"
)


func main()  {
	u, err := user.Current()
	if err != nil {
		fmt.Println("Couldn't find user")
	}

	fmt.Println(u.Username)
}
