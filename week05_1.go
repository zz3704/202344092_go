package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("정수 입력 : ")
	reader := bufio.NowReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n')
	fmt.Println(inputScore)
}
