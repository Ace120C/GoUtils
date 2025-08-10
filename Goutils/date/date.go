package main

import (
	"fmt"
	"time"
)


func main()  {
	Now := time.Now()

	layout := "Mon Jan 2 03:04:05 PM -0700 2006"

	fmt.Println(Now.Format(layout))
}
