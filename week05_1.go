package main

import (
	"fmt"
	"math/rand"
)

func main() {
	dice := rand.Intn(6) + 1 //0~5 add 1 to make dice number(1~6)
	fmt.Println(dice)
}
