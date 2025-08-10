package main

import (
	"fmt"
	"time"
)


func main()  {
	Now := time.Now()

	Date := Now.Format(time.UnixDate)

	fmt.Println(Date)
}
