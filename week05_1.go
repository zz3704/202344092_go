package main

import (
	"time"
	"fmt"
)

func main()  {
	var now time.Time=time.now()
	var year int = now.year()
	month := now.Month()
	fmt.println(now.Month)
	fmt.println(year)
}