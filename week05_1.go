package main

import()

func main()  {
	var now time.Time=time.now()
	var year int = now.year()
	fmt.println(year)
}